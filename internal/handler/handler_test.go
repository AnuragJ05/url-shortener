package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
	"url-shortener/constants"
	"url-shortener/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	// Test case: Valid URLStore and mutex
	URLStore := &models.URLStore{
		URLMap:    make(map[string]string),
		DomainMap: make(map[string]int),
	}
	mu := &sync.Mutex{}

	go Handle(URLStore, mu)

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Test case: Shorten a URL
	response, err := http.Post("http://localhost:8080/shorten", "application/json", bytes.NewBufferString(`{"url": "https://example.com"}`))
	require.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result struct {
		URL string `json:"url"`
	}
	err = json.NewDecoder(response.Body).Decode(&result)
	require.NoError(t, err)

	shortURL := result.URL
	// Assuming result.URL returns the full URL like "http://localhost:8080/8xmdYWws"

	// Extract the short identifier from the full short URL
	shortIdentifier := shortURL[len(constants.BaseURL):]

	// Test case: Redirect to the original URL
	response, err = http.Get(fmt.Sprintf("http://localhost:8080/%s", shortIdentifier))
	require.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Test case: Get the top visited domains
	response, err = http.Get("http://localhost:8080/topdomains?count=3")
	require.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}
