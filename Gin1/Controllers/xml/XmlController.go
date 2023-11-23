package xml

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Desc    string `json:"desc" xml:"desc"`
	Content string `json:"content" xml:"content"`
}

func Xml1(c *gin.Context) {

	c.XML(http.StatusOK, gin.H{
		"success": true,
		"message": "hello gin",
	})

}

func Xml2(c *gin.Context) {
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

}
