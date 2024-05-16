
package AdminURIHandlers

import (
	"github.com/gin-gonic/gin"
)

func AdminHomeUriGetLoginMethodHandler(c *gin.Context) {
	
	c.JSON(
		200,
		gin.H{
			"status": "Admin panel",
		},
	)

}

