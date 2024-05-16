
package Services


import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/Nabin-Flash320/go_phonebook/App/Phonebook/URIHandlers"
)

type AppURIS struct {
	method string
	uri string
	handler gin.HandlerFunc
}

type URIPatterns struct {
	app string // This will be used for grouping the routes i.e. <appe>/<uri>
	appuris []AppURIS // All the URIs relating to <app> will be held by this array
}

var uri_patterns = [...] URIPatterns {
	{
		app: "phonebook",
		appuris: []AppURIS{
			{
				method: "GET",
				uri: "/get/records",
				handler: URIHandlers.HomeUriGetMethodHandler,
			},
			{
				method: "POST",
				uri: "/set/records",
				handler: URIHandlers.HomeUriPostMethodHandler,
			},
			{
				method: "GET",
				uri: "/get/:id",
				handler: URIHandlers.HomeUriGetByIDMethodHandler,
			},
			{
				method: "POST",
				uri: "/del/:id",
				handler: URIHandlers.HomeUriPostDeleteRecordMethodHandler,
			},
		},
	},
}

func ServerRegisterUIR(engine *gin.Engine) {

	for _, uri_pattern := range uri_patterns {

		if uri_pattern.app != "" {

			for _, appuri := range uri_pattern.appuris {

				group_name := fmt.Sprintf("api/%s", uri_pattern.app)
				group := engine.Group(group_name)
				{
					
					method := fmt.Sprintf("%s", appuri.uri)
					if appuri.method == "GET" {
						
						group.GET(method, appuri.handler)

					} else if appuri.method == "POST" {

						group.POST(method, appuri.handler)

					}

				}

			}

		}

	}

}







