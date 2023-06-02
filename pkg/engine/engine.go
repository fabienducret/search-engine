package engine

import (
	"searchengine/pkg/async"
	"searchengine/pkg/entities"
)

type Engine struct {
	providers []Provider
}

func NewEngine(providers []Provider) Engine {
	return Engine{providers: providers}
}

func (e Engine) Search(query string) []entities.SearchResult {
	searchResults := e.resultsFromProviders(query)

	return removeDuplicateLinks(searchResults)
}

func (e Engine) resultsFromProviders(query string) []entities.SearchResult {
	var searchResults []entities.SearchResult

	async.Range(e.providers, func(p Provider) {
		resultsForProvider := p.Search(query)
		searchResults = append(searchResults, resultsForProvider...)
	})

	return searchResults
}

func removeDuplicateLinks(toDeduplicate []entities.SearchResult) []entities.SearchResult {
	var searchResults []entities.SearchResult

	for _, s := range toDeduplicate {
		if !contains(searchResults, s.Link) {
			searchResults = append(searchResults, s)
		}
	}

	return searchResults
}

func contains(searchResults []entities.SearchResult, link string) bool {
	for _, s := range searchResults {
		if s.Link == link {
			return true
		}
	}

	return false
}
