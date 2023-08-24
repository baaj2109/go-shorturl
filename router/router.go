package router

import (
	"github.com/baaj2109/shorturl/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())
	// router.Use(middlewares.Request()).StaticFile("/", "./public").GET("/:code", controller.Path)
	v1 := router.Group("/v1")
	v1.Use(middlewares.Request())
	// v1.POST("/create", controller.Create)
	// v1.POST("/multicreate", controller.multicreate)
	// v1.POST("/query", controller.Query)
	return router
}
