package router

import (
	"mangosteen/config"
	"mangosteen/internal/controller"
	"mangosteen/internal/database"

	_ "mangosteen/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func New() *gin.Engine {
	config.LoadAppConfig()
	database.Connect()
	r := gin.Default()

	r.GET("/api/v1/ping", controller.Ping)
	r.POST("/api/v1/validation_code", controller.CreateValidationCode)
	r.POST("/api/v1/session", controller.CreateSession)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
