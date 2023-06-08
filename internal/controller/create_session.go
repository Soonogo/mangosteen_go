package controller

import (
	"mangosteen/config/queries"
	"mangosteen/internal/database"

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
	jwt := "xxx"
	respBody := struct {
		Jwt string `json:"jwt"`
	}{
		Jwt: jwt,
	}
	c.JSON(200, respBody)
}
