package viewmodel

import (
	"html/template"
	"searchengine/pkg/entities"
)

type searchResult struct {
	Title       string
	Description template.HTML
	Link        string
	From        string
}

func Response(query string, searchResults []entities.SearchResult) map[string]interface{} {
	var results []searchResult

	for _, res := range searchResults {
		results = append(results, searchResult{
			Title:       res.Title,
			Link:        res.Link,
			Description: template.HTML(res.Description),
			From:        res.From,
		})
	}

	return map[string]interface{}{"query": query, "results": results}
}
