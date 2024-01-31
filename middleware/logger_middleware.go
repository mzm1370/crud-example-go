package middleware

import (
	"github.com/gin-gonic/gin"

	"log"
)

func LoggerMiddleware() gin.HandleFunc {
	return func(c *gin.Context) {
		log.Println("Request received")
		c.Next()
	}
}
