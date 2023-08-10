package router

import (
	"github.com/baaj2109/shorturl/controller.go"
	"github.com/baaj2109/shorturl/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AddAllowHeaders("Authorization", "Content-Type")
	config.AddAllowMethods("GET", "POST")
	router.Use(cors.New(config))

	versionRouter := router.Group("/v1")
	versionRouter.GET("/:url_id/", middleware.Cache(), controller.GetUrlHandler)
	versionRouter.POST("/urls/", controller.CreateUrlHandler)

	return router
}
