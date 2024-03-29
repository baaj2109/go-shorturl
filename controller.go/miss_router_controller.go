package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MissingRouteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":       "This routing path has not supported yet, please check out the API documentation below.",
		"documentation": "https://documenter.getpostman.com/view/12176709/UVypycK7",
	})
	return
}
