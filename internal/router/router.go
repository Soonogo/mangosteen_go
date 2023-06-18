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

func loadController() []controller.Controller {
	return []controller.Controller{
		&controller.SessionController{},
		&controller.ValidationCodeController{},
		&controller.MeController{},
	}
}
func New() *gin.Engine {
	config.LoadAppConfig()
	database.Connect()
	r := gin.Default()

	api := r.Group("/api")

	for _, c := range loadController() {
		c.RegisterRoutes(api)
	}

	r.GET("/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
