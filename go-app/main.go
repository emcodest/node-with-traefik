package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	portConnect := os.Getenv("PORT")
	// Initialize Gin
	router := gin.Default()
	// Use CORS middleware
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		// make http get request
		res, err := http.Get("http://node-server:3111")
		defer res.Body.Close()
		if err != nil {

			c.String(400, "Error occured ...")
			return
		}
		str, _ := io.ReadAll(res.Body)
		str2 := fmt.Sprintf("Traefik: %v", string(str))
		c.String(200, str2)

		//c.JSON(200, gin.H{"message": "FlipNaira API"})
	})
	// health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	// Start the server
	fmt.Println("Server started at http://localhost:" + portConnect)
	router.Run(":" + portConnect)

}
