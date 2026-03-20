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

	// serve html
	r.StaticFile("/", "./index.html")
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.POST("/create-short-url", handler.CreateShortUrl)

	r.GET("/:shortUrl", handler.HandleShortUrlRedirect)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Running on port:", port)

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}