package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

const NumLines = 123866

type Quote struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Name string `json:"name"`
}

func main() {
	// Initialize a new Gin router
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	router.GET("/", func(c *gin.Context) {
		f, _ := os.Open("quotes.json.lines")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for i := 0; i < rand.Intn(NumLines-1); i++ {
			scanner.Scan()
		}
		var line string
		for scanner.Scan() {
			line = scanner.Text()
			break
		}
		log.Println(line)
		var q Quote
		json.Unmarshal([]byte(line), &q)
		c.JSONP(200, q)
	})
	router.Run(":8063")
}
