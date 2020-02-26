package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Commom() gin.HandlerFunc {
	return func(c *gin.Context) {
		GetFile(c)
		c.Next()
	}
}
