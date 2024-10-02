package shortener

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"url-shortener/constants"
	"url-shortener/internal/utils"
	"url-shortener/models"

	"github.com/gin-gonic/gin"
)

// Shorten is a function that handles the shortening of URLs.
func Shorten(URLStore *models.URLStore, mu *sync.Mutex) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.URLRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		originalURL := request.URL
		fmt.Println("Original URL: ", originalURL)
		if _, err := url.ParseRequestURI(originalURL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		}
		mu.Lock()

		for shortURL, longURL := range URLStore.URLMap {
			if longURL == originalURL {
				c.JSON(http.StatusOK, gin.H{"url": shortURL})
				return
			}
		}
		shortURL := constants.BaseURL + utils.GenerateRandomString(8)
		URLStore.URLMap[shortURL] = originalURL
		c.JSON(http.StatusOK, gin.H{"url": shortURL})
		fmt.Println("Updated URLMap: ", URLStore.URLMap)
		mu.Unlock()
	}
}

func RedirectURL(URLStore *models.URLStore, mu *sync.Mutex) gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		shortURL := constants.BaseURL + c.Param("short")
		longURL := ""
		for short, long := range URLStore.URLMap {
			if short == shortURL {
				longURL = long
				break
			}
		}
		if longURL == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}
		fmt.Println("longURL: ", longURL)
		c.Redirect(http.StatusFound, longURL)
	}
}
