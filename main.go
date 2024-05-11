
package main



import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/Nabin-Flash320/go_phonebook/URIHandlers"
	"github.com/Nabin-Flash320/go_phonebook/Middlewares"
	"github.com/Nabin-Flash320/go_phonebook/ModelDB"
	"github.com/Nabin-Flash320/go_phonebook/Settings"
)

func serverInit() {

	host_addr := Settings.HostSettingsGetValidHostAddress()
	
	router := gin.Default()

	router.Use(Middlewares.LoggerwareMiddleware())

	router.GET("/", URIHandlers.HomeUriGetMethodHandler)
	
	fmt.Printf("\033[32m Server started at: http://%s:8000 \n\033[0m", host_addr)

	router.Run(host_addr + ":8000")


}


func main() {

	var serve *bool = flag.Bool("serve", false, "Run the server")
	var migrate *bool = flag.Bool("migrate", false, "Migrate models to the database")
	var model *string = flag.String("model", "all", "Migrate specific model to the database")

	flag.Parse()

	if *serve {

		serverInit()

	}else if *migrate {

		ModelDB.ModelDBMakeMigrations(*model)

	}


}


