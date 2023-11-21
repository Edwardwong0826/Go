package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Desc    string `json:"desc" xml:"desc"`
	Content string `json:"content" xml:"content"`
}

type Article2 struct {
	// `` this symbol called backquote 反引号, for json body key to allow to have small letter
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// use below command to build and run at one time, else just run it will prompt firewall everytime
// go build main.go && .\main.exe

// for hot reload, there is two library, use the first one
// once install air, just run air
// air - https://github.com/cosmtrek/air
// gravityblast - https://github.com/gravityblast/fresh

// currently air hot reload not support well debugger
// when debug in VS code, just go click left section RUN AND DEBUG, click this Launch Package button
// https://superuser.com/questions/1785007/can-i-safely-disable-a-firewall-rule-added-by-visual-studio-code
// if keep block by windows firewall, go to Windows Defender Firewall with Advanced Security
// go to Inbound Rules -> New Rule -> Port -> put 8080 port -> and next

// https://maelvls.dev/go111module-everywhere/
// when get if there is error about GO111MODULE=off
// in windows - set GO111MODULE=on
// so that can use go install xxx
func main() {
	// create default router engine
	r := gin.Default()

	// config router
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "This is : %v", "hello gin")
	})

	r.GET("/user", func(c *gin.Context) {

		// Query to get parameter
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	r.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")

		c.JSON(http.StatusOK, gin.H{
			"message": "news info",
			"id":      id,
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "News page")
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "hello gin",
		})

	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "hello gin",
		})

	})

	r.GET("/json3", func(c *gin.Context) {

		// dont know why, use this way to display JSON, type fields first letter required to be capital letter
		a := &Article2{
			Title:   "I am a title",
			Desc:    "description",
			Content: "test content",
		}

		c.JSON(http.StatusOK, a)

	})

	// JSONP is for cross origin，CORS to handle JSONP request
	// it can put callback on the endpoint example
	// localhost:8080/jsonp?callback=xxx
	r.GET("/jsonp", func(c *gin.Context) {

		a := &Article{
			Title:   "I am a title",
			Desc:    "description",
			Content: "test content",
		}

		c.JSONP(http.StatusOK, a)

	})

	r.GET("/xml", func(c *gin.Context) {

		c.XML(http.StatusOK, gin.H{
			"success": true,
			"message": "hello gin",
		})

	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post request to create data")
	})

	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}

		// c.ShouldBind when try run postman get with JSON request body, it JSON response body will empty
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.POST("/addUser", func(c *gin.Context) {

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
	})

	// get POST XML data
	// for payment processing, is possible to pass XML data back
	r.POST("/xml", func(c *gin.Context) {
		article := &Article{}

		xmlSlice, _ := c.GetRawData()
		fmt.Println(article)
		if err := xml.Unmarshal(xmlSlice, &article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "edit request to edit data")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "delete request to delete data")
	})

	// reuqest param for router

	r.GET("/list/:cid", func(c *gin.Context) {

		param := c.Param("cid")
		c.String(http.StatusOK, param)

	})
	// start a web server
	// default start HTTP server at 0.0.0.0:8080 port
	r.Run(":8080")
}
