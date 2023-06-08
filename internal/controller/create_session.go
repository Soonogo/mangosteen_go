package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
	var reqBody struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	q := database.NewQuery()
	_, err := q.FindValidationCode(c, queries.FindValidationCodeParams{
		Email: reqBody.Email,
		Code:  reqBody.Code,
	})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := q.FindUserByEmail(c, reqBody.Email)

	if err != nil {
		user, err = q.CreateUser(c, reqBody.Email)
		if err != nil {
			log.Println("CreateUser fail", err)
			c.String(http.StatusInternalServerError, "wait a minute")
			return
		}
	}

	jwt, err := jwt_helper.GenerateJWT(int(user.ID))
	if err != nil {
		log.Println("GenerateJWT fail", err)
		c.String(http.StatusInternalServerError, "wait a minute")
		return
	}

	respBody := gin.H{
		"jwt":    jwt,
		"userId": user.ID,
	}
	c.JSON(200, respBody)
}
