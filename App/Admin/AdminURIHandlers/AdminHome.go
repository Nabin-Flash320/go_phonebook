
package AdminURIHandlers

import (
	"github.com/gin-gonic/gin"
)

func AdminHomeUriPostLoginMethodHandler(c *gin.Context) {

	c.JSON(
		200,
		gin.H{
			"status": "success",
		},
	)

}

func AdminHomeUriPostLogoutMethodHandler(c *gin.Context) {

	c.JSON(
		200,
		gin.H{
			"status": "succcess"
		},
	)

}

