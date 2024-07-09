package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func testePostA(c *gin.Context) {
	log.Println("Called POST A")
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Print(err)
	}
	log.Println(string(jsonData))
	c.JSON(200, string(jsonData))
}

func testePostB(c *gin.Context) {
	log.Println("Called POST B")
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Print(err)
	}
	log.Println(string(jsonData))
	c.JSON(200, string(jsonData))
}

func main() {
	r := gin.New()

	r.POST("/testepostA", testePostA)
	r.POST("/testepostB", testePostB)

	r.Run(":8181")
}
