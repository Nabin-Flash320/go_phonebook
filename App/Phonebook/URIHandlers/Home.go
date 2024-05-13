
package URIHandlers

import (
	"github.com/gin-gonic/gin"
)



func HomeUriGetMethodHandler(c *gin.Context) {

	c.JSON(
		200, 
		gin.H{
			"status": "success",
		},
	)

}



