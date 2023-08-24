package controller

import (
	"net/http"

	"github.com/baaj2109/shorturl/common"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	// userId := c.GetInt("userId")
	lUrl := c.PostForm("url")
	common.SugarLogger.Info("incomming create url request, url:" + lUrl)
	if len(lUrl) == 0 {
		common.SugarLogger.Info("Url is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "url is empty",
		})
		return
	}

}
