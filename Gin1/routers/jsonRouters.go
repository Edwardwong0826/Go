package routers

import (
	"Go/Gin1/Controllers/json"

	"github.com/gin-gonic/gin"
)

// if want this function can be use outside, this function should define as public
// where the first letter of function should be capital letter
func JsonRoutersInit(r *gin.Engine) {

	jsonRouters := r.Group("/json")
	{
		jsonRouters.GET("/json1", json.Json1)

		jsonRouters.GET("/json2", json.Json2)

		jsonRouters.GET("/json3", json.Json3)

		// JSONP is for cross origin，CORS to handle JSONP request
		// it can put callback on the endpoint example
		// localhost:8080/jsonp?callback=xxx
		jsonRouters.GET("/jsonp", json.Jsonp)
	}

}
