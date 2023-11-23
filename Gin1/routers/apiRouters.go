package routers

import (
	"Go/Gin1/Controllers/api"

	"github.com/gin-gonic/gin"
)

// if want this function can be use outside, this function should define as public
// where the first letter of function should be capital letter
func ApiRoutersInit(r *gin.Engine) {

	apiRouters := r.Group("/api")
	{

		// config router
		// notice GetUser without () is not function
		// here we are register GetUser
		apiRouters.GET("/getUser", api.ApiController{}.User)

		apiRouters.POST("/addUser", api.ApiController{}.Add)

		apiRouters.PUT("/edit", api.ApiController{}.Put)

		apiRouters.DELETE("/delete", api.ApiController{}.Delete)

		// reuqest param for router
		apiRouters.GET("/list/:cid", api.ApiController{}.List)
	}
}
