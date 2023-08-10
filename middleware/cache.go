package middleware

import (
	"net/http"

	"github.com/baaj2109/shorturl/constants"
	"github.com/baaj2109/shorturl/database"
	"github.com/gin-gonic/gin"
)

func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlId := c.Param("url_id")
		if len(urlId) == 0 {
			c.JSON(constants.URL_ID_MISSING_ERROR.StatusCode, constants.URL_ID_MISSING_ERROR)
			c.Abort()
			return
		}

		redisClient := database.DatabaseClient.RedisClient

		idContent, err := redisClient.Get(urlId).Result()
		if err != nil {
			// not in cache or expire
			c.Next()
			return
		}
		if idContent == constants.ID_NOT_EXIST {
			c.JSON(constants.URL_ID_NOT_EXIST_ERROR.StatusCode, constants.URL_ID_NOT_EXIST_ERROR)
			c.Abort()
			return
		}

		// this url content is in cache
		c.Redirect(http.StatusFound, idContent)
		c.Abort()
		return

	}
}
