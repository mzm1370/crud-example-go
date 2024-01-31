package handlers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LogInputMiddleware(c *gin.Context) {
	log.Printf("[%s] %s - %s", c.Request.Method, c.Request.URL.Path, time.Now().Format("2006-01-02 15:04:05"))

	log.Println("Headers:")
	for key, values := range c.Request.Header {
		log.Printf("%s: %v", key, values)
	}

	if c.Request.ContentLength > 0 {
		body, err := c.GetRawData()
		if err != nil {
			log.Printf("Error reading request body: %v", err)
		} else {
			log.Printf("Request Body: %s", body)
			// c.Request.Body = gin.NopCloser(c.Request.Body)
		}
	}

	c.Next()

}
