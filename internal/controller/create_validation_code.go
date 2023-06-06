package controller

import (
	"log"

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
		Email string
	}
	c.ShouldBindJSON(&body)
	log.Println("------------")
	log.Println(body)
	c.String(200, "ok")
}
