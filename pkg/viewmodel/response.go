package viewmodel

import (
	"html/template"
	"searchengine/pkg/domain"
)

type searchResult struct {
	Title       string
	Description template.HTML
	Link        string
}

func Response(query string, searchResults []domain.SearchResult) map[string]interface{} {
	var results []searchResult

	for _, res := range searchResults {
		results = append(results, searchResult{
			Title:       res.Title,
			Link:        res.Link,
			Description: template.HTML(res.Description),
		})
	}

	return map[string]interface{}{"query": query, "results": results}
}
