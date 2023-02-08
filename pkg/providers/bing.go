package providers

import (
	"searchengine/pkg/domain"
	"searchengine/pkg/net"
)

type Bing struct {
	HttpClient net.HttpClient
}

func (p Bing) Search(q string) []domain.SearchResult {
	return []domain.SearchResult{{
		Title:       "Fake title",
		Description: "Fake description",
		Link:        "www.google.com",
	}}
}
