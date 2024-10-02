package main

import (
	"sync"
	"url-shortener/internal/handler"
	"url-shortener/models"
)

var mu sync.Mutex

func main() {

	urlStore := models.URLStore{
		URLMap:    map[string]string{},
		DomainMap: map[string]int{},
	}

	handler.Handle(&urlStore, &mu)

}
