package shortener

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
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
		defer mu.Unlock()

		for shortURL, longURL := range URLStore.URLMap {
			if longURL == originalURL {
				c.JSON(http.StatusOK, gin.H{"url": shortURL})
				return
			}
		}

		domain, _ := url.Parse(originalURL)
		if _, ok := URLStore.DomainMap[domain.Host]; !ok {
			URLStore.DomainMap[domain.Host] = 1
		} else {
			URLStore.DomainMap[domain.Host] += 1
		}

		shortURL := constants.BaseURL + utils.GenerateRandomString(8)
		URLStore.URLMap[shortURL] = originalURL

		c.JSON(http.StatusOK, gin.H{"url": shortURL})
		fmt.Println("Updated URLMap: ", URLStore.URLMap)
		fmt.Println("Updated DomainMap: ", URLStore.DomainMap)
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

func GetMostVisitedDomainsHandler(URLStore *models.URLStore, mu *sync.Mutex) gin.HandlerFunc {
	return func(c *gin.Context) {
		countStr := c.Query("count")
		count, err := strconv.Atoi(countStr)

		if err != nil || count <= 0 {
			count = 3
		}
		top := GetMostVisitedDomains(URLStore, mu, count)
		c.JSON(http.StatusOK, top)
	}
}

func GetMostVisitedDomains(URLStore *models.URLStore, mu *sync.Mutex, count int) []models.KV {
	mu.Lock()
	defer mu.Unlock()

	// Create the slice
	var kvSlice []models.KV
	for k, v := range URLStore.DomainMap {
		kvSlice = append(kvSlice, models.KV{Key: k, Value: v})
	}

	// Sort the slice by value (occurrence count) in descending order
	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].Value > kvSlice[j].Value
	})

	// Get the top keys with the highest occurrences
	top3 := kvSlice
	if len(kvSlice) > count {
		top3 = kvSlice[:count]
	}

	return top3
}
