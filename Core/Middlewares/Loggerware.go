package Middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)



func LoggerwareMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		duration := time.Since(start)

		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)

	}

}



