
package AdminURIHandlers

import (
	"github.com/gin-gonic/gin"
)

func AdminHomeUriGetLoginMethodHandler(c *gin.Context) {

	c.HTML(
		200,
		"AdminLogin.tmpl",
		gin.H{
			"title": "Admin Login",
		},
	)

}

