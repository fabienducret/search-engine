package providers

import "searchengine/pkg/domain"

type Google struct {
}

func (p Google) Search(q string) []domain.SearchResult {
	return []domain.SearchResult{{
		Title:       "Title 1 de google",
		Description: "Ceci est la première description",
		Link:        "https://google.com",
	}, {
		Title:       "Title 2 de google",
		Description: "Ceci est la deuxième description",
		Link:        "https://google.com",
	}}
}
