package routers

import (
	"Go/Gin1/Controllers/origin"

	"github.com/gin-gonic/gin"
)

// if want this function can be use outside, this function should define as public
// where the first letter of function should be capital letter
func OriginRoutersInit(r *gin.Engine) {

	originRouters := r.Group("/")
	{
		// config router
		originRouters.GET("/get", origin.Get)

		originRouters.POST("/post", origin.Post)

		originRouters.GET("/user", origin.User)

		originRouters.GET("/article", origin.Article)

		originRouters.GET("/news", origin.News)

	}
}
