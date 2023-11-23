package origin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.String(http.StatusOK, "This is : %v", "hello gin")
}

func Post(c *gin.Context) {
	c.String(200, "post request to create data")
}

func User(c *gin.Context) {

	// Query to get parameter
	username := c.Query("username")
	age := c.Query("age")
	page := c.DefaultQuery("page", "1")

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"age":      age,
		"page":     page,
	})
}

func Article(c *gin.Context) {
	id := c.DefaultQuery("id", "1")

	c.JSON(http.StatusOK, gin.H{
		"message": "news info",
		"id":      id,
	})
}

func News(c *gin.Context) {
	c.String(200, "News page")
}
