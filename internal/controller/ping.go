package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc

// @Summary		Show an account
// @Description	get string by ID
// @Tags			ping
// @Accept			json
// @Produce		json
// @Success		200
// @Failure		500
// @Router			/ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
