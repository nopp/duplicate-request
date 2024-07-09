package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func duplicateRequest(jsonData []byte, url string) string {
	jsonStr := []byte(jsonData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Print(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err.Error())
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func main() {
	r := gin.New()

	r.POST("/middleware", func(c *gin.Context) {
		jsonData, err := c.GetRawData()
		if err != nil {
			log.Print(err)
		}
		// Environment A
		log.Print("Send to environment A")
		log.Print(duplicateRequest(jsonData, "http://localhost:8181/testepostA"))
		// Environment B
		log.Print("Send to environment B")
		log.Print(duplicateRequest(jsonData, "http://localhost:8181/testepostB"))
		c.String(200, "Sented")
	})

	log.Print("Duplicator started")
	r.Run(":8080")
}
