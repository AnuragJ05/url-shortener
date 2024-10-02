package models

type URLStore struct {
	URLMap    map[string]string `json:"url"`
	DomainMap map[string]int    `json:"domain"`
}

type URLRequest struct {
	URL string `json:"url" binding:"required"`
}

type KV struct {
	Key   string
	Value int
}
