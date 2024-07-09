package cmd

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

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

func Execute() {

	envs, exists := os.LookupEnv("ENVIRONMENTS")
	if !exists {
		log.Panic("You need to export env ENVIRONMENTS")
	}

	r := gin.New()

	r.POST("/middleware", func(c *gin.Context) {
		jsonData, err := c.GetRawData()
		if err != nil {
			log.Print(err)
		}
		for _, envUrl := range strings.Split(envs, ";") {
			log.Print("Send to %V", envUrl)
			log.Print(duplicateRequest(jsonData, envUrl))
		}
		c.String(200, "Sented")
	})

	log.Print("Duplicator started")
	r.Run(":8080")
}
