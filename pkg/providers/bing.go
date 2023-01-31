package providers

import "searchengine/pkg/domain"

type Bing struct {
}

func (p Bing) Search(q string) []domain.SearchResult {
	return []domain.SearchResult{{
		Title:       "Title 1 de bing",
		Description: "Ceci est la première description",
		Link:        "https://google.com",
	}, {
		Title:       "Title 2 de bing",
		Description: "Ceci est la deuxième description",
		Link:        "https://google.com",
	}}
}
