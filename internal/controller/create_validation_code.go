package controller

import (
	"log"
	"mangosteen/internal/email"

	"github.com/gin-gonic/gin"
)

// validation godoc
// @Summary      validation code
// @Description  post validation code
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /createValidationCode [post]
func CreateValidationCode(c *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(400, "参数错误")
		return
	}
	if err := email.SendValidationCode(body.Email, "123456"); err != nil {
		log.Println("[SendValidationCode fail]", err)
		c.String(500, "发送失败")
		return
	}
	c.Status(200)
}
