package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
)

func mainBlock() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.Run(":8061")
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandlerXML)

	router.Run(":8060")
}

func IndexHandlerJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "baki was here !...",
	})
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func IndexHandlerXML(c *gin.Context) {
	c.XML(200, Person{FirstName: "Baki",
		LastName: "Sevgul"})
}
