package models

type URLStore struct {
	URLMap map[string]string `json:"url"`
}

type URLRequest struct {
	URL string `json:"url" binding:"required"`
}
