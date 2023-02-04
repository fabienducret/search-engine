package providers

import (
	"searchengine/pkg/domain"
)

type Bing struct {
}

func (p Bing) Search(q string) []domain.SearchResult {
	return []domain.SearchResult{{
		Title:       "toto",
		Description: "ddd",
		Link:        "www.google.com",
	}}
}
