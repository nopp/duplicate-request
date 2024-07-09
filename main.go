package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func duplicatePost(jsonData []byte, url string) {
	jsonStr := []byte(jsonData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	r := gin.New()

	r.POST("/btgmiddleware", func(c *gin.Context) {
		jsonData, err := c.GetRawData()
		if err != nil {
			fmt.Println("Cannot retrieve data")
		}
		// Environment A
		duplicatePost(jsonData, "http://localhost:8181/testepostA")
		// Environment B
		duplicatePost(jsonData, "http://localhost:8181/testepostB")
	})

	r.Run(":8080")
}
