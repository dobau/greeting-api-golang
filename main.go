package main

import (
	"net/http"
	"os"

	"github.com/dobau/greeting-api-golang/rest"
	"github.com/gin-gonic/gin"
)

type Greeting struct {
	Greeting string `json:greeting`
}

func main() {
	r := gin.Default()
	r.GET("/greeting", addCors, func(c *gin.Context) {
		// Create a Resty Client
		rest := rest.New()

		greeting := new(Greeting)
		err := rest.Get("https://workshop-go-greeting.herokuapp.com/greet", http.Header{}, greeting)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(200, gin.H{
			"owner":      "Rafael Alves",
			"greeting":   greeting.Greeting,
			"repository": "https://github.com/dobau/greeting-api-golang",
		})
	})

	r.Run(port()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func addCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Authorization, X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Allow-Request-Method")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	c.Header("Allow", "GET, POST, OPTIONS, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func port() string {
	// heroku variable
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}
