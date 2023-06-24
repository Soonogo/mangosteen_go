package internal

import (
	"mangosteen/config"
	"mangosteen/internal/database"
	"mangosteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	config.LoadAppConfig()
	database.Connect()
	r.Use(middleware.Me([]string{"/api/v1/validation_codes", "/api/v1/session", "/ping"}))
}
