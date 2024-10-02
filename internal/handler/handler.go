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

	r.Run(":8080")
}
