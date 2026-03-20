package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/YamamoJuan/go-url-shortener/shortener"
	"github.com/YamamoJuan/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func normalizeUrl(url string) string {
	if !strings.HasPrefix(url, "http://") &&
		!strings.HasPrefix(url, "https://") {

		url = "https://" + url
	}
	return url
}

func CreateShortUrl(c *gin.Context) {

	var creationRequest UrlCreationRequest

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	longUrl := normalizeUrl(creationRequest.LongUrl)

	shortUrl := shortener.GenerateShortLink(
		longUrl,
		creationRequest.UserId,
	)

	store.SaveUrlMapping(
		shortUrl,
		longUrl,
		creationRequest.UserId,
	)

	host := os.Getenv("BASE_URL")

	c.JSON(200, gin.H{
		"message": "short url created succesfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}