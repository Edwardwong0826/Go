package admin

import (
	"Go/Gin2/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (con IndexController) Index(c *gin.Context) {

	// this c.Get can get the middleware set value key username
	username, _ := c.Get("username")
	fmt.Println(username)
	fmt.Println(models.UnixToTime(1629788564))
	fmt.Println("----------------------------")
	// cast to string
	v, ok := username.(string)
	if ok {
		c.String(200, "User list---"+v)

	} else {
		c.String(200, "User list---get user fail")
	}

}
