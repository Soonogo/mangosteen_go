package middleware

import (
	"errors"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Me(whiteList []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		index := indexOf(whiteList, path)

		if index != -1 {
			c.Next()
			return
		}
		user, err := getMe(c)
		if err != nil {
			c.String(401, "unauthorized")
			return
		}
		c.Set("me", user)
		c.Next()
	}
}
func getMe(c *gin.Context) (queries.User, error) {
	var user queries.User
	auth := c.GetHeader("Authorization")
	if len(auth) < 8 {
		c.String(401, "unauthorized")
		return user, errors.New("unauthorized")
	}
	jwtString := auth[7:]
	t, err := jwt_helper.Parse(jwtString)
	if err != nil {
		c.String(401, "unauthorized")
		return user, errors.New("unauthorized")
	}
	m, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		c.String(401, "unauthorized")
		return user, errors.New("unauthorized")
	}
	userID, ok := m["user_id"].(float64)
	if !ok {
		c.String(401, "unauthorized")
		return user, errors.New("unauthorized")

	}
	userIDInt := int32(userID)
	if err != nil {
		c.String(401, "无效的JWT")
		return user, errors.New("unauthorized")
	}
	q := database.NewQuery()
	userIDInt32 := int32(userIDInt)
	user, err = q.FindUser(c, userIDInt32)
	if err != nil {
		c.String(401, "无效的JWT")
		return user, errors.New("unauthorized")
	}
	return user, nil
}

func indexOf(slice []string, s string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}
	return -1
}
