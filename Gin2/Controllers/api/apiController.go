package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type ApiController struct {
	BaseController
}

// this (con xxxController) allow all function inherited from Apicontroller struct
// so that in routers, it register like <folder name>.<type name>.<function name>
// similiar to Java OOP class method name
func (con ApiController) User(c *gin.Context) {
	user := &UserInfo{}

	// c.ShouldBind when try run postman get with JSON request body, it JSON response body will empty
	// it should be other format
	if err := c.ShouldBind(&user); err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
	con.Success(c)
}

func (con ApiController) Add(c *gin.Context) {

	// when use in postman, use c.BindJSON only the payload can display JSON values
	var requestBody User
	if err := c.BindJSON(&requestBody); err != nil {

	}

	// username := c.PostForm("username")
	// password := c.PostForm("password")

	c.JSON(http.StatusOK, gin.H{
		"usernamee": requestBody.Username,
		"password":  requestBody.Password,
	})
}

func (con ApiController) Put(c *gin.Context) {
	c.String(200, "edit request to edit data")
}

func (con ApiController) Delete(c *gin.Context) {
	c.String(200, "delete request to delete data")
}

func (con ApiController) List(c *gin.Context) {

	param := c.Param("cid")
	c.String(http.StatusOK, param)

}
