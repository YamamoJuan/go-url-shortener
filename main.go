package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/YamamoJuan/go-url-shortener/handler"
	"github.com/YamamoJuan/go-url-shortener/store"
)

func main() {

	store.InitializeStore()

	r := gin.Default()
	r.Use(cors.Default())

	r.StaticFile("/", "./index.html")
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API!",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Running on port:", port)

	err := r.Run(":" + port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}