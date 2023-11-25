package routers

import (
	"Go/Gin2/Controllers/admin"
	"Go/Gin2/middleware"

	"github.com/gin-gonic/gin"
)

// if want this function can be use outside, this function should define as public
// where the first letter of function should be capital letter
func AdminRoutersInit(r *gin.Engine) {

	adminRouters := r.Group("/admin", middleware.InitMiddleware)
	{

		adminRouters.GET("/get", admin.Get)

		adminRouters.POST("/post", admin.Post)

		adminRouters.GET("/user", admin.User)

		adminRouters.GET("/article", admin.Article)

		adminRouters.GET("/news", admin.News)

		adminRouters.GET("/index", admin.IndexController{}.Index)

	}
}
