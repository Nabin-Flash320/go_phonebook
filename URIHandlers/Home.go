
package URIHandlers

import (
	"github.com/gin-gonic/gin"
)



func HomeUriGetMethodHandler(c *gin.Context) {

	c.String(200, "Hello, World: from home get uri handler")

}



