package routers

import (
	"Go/Gin1/Controllers/xml"

	"github.com/gin-gonic/gin"
)

// if want this function can be use outside, this function should define as public
// where the first letter of function should be capital letter
func XmlRoutersInit(r *gin.Engine) {

	xmlRouters := r.Group("/xml")
	{

		xmlRouters.GET("/xml1", xml.Xml1)

		// get POST XML data
		// for payment processing, is possible to pass XML data back
		xmlRouters.POST("/xml2", xml.Xml2)

	}

}
