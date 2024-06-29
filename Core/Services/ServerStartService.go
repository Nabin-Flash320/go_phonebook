
package Services

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/Nabin-Flash320/go_phonebook/Core/Settings"
	"github.com/Nabin-Flash320/go_phonebook/Core/Middlewares"
)



func ServerInit() {

	host_addr := Settings.HostSettingsGetValidHostAddress()
	
	router := gin.Default()

	router.Use(Middlewares.LoggerwareMiddleware())

	ServerRegisterUIR(router)
	
	fmt.Printf("\033[32m Server started at: http://%s:8000 \n\033[0m", host_addr)

	router.Run(host_addr + ":8000")

}


