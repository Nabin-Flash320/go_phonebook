
package main



import (
	"github.com/gin-gonic/gin"

	"github.com/Nabin-Flash320/go_phonebook/URIHandlers"
	"github.com/Nabin-Flash320/go_phonebook/Middlewares"
)




func main() {
	
	router := gin.Default()

	router.Use(Middlewares.LoggerwareMiddleware())

	router.GET("/", URIHandlers.HomeUriGetMethodHandler)

	router.Run("127.0.0.1:8000")

}


