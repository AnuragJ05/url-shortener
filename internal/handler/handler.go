package handler

import (
	"sync"
	"url-shortener/internal/shortener"
	"url-shortener/models"

	"github.com/gin-gonic/gin"
)

func Handle(URLStore *models.URLStore, mu *sync.Mutex) {

	r := gin.Default()

	r.POST("/shorten", shortener.Shorten(URLStore, mu))

	r.GET("/:short", shortener.RedirectURL(URLStore, mu))

	r.GET("/topdomains", shortener.GetMostVisitedDomainsHandler(URLStore, mu))

	r.Run(":8080")
}
