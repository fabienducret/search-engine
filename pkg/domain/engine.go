package domain

import "searchengine/pkg/async"

type Engine struct {
	providers []Provider
}

func NewEngine(providers []Provider) Engine {
	return Engine{providers: providers}
}

func (e Engine) Search(query string) []SearchResult {
	searchResults := e.resultsFromProviders(query)

	// TODO tri, fusion

	return searchResults
}

func (e Engine) resultsFromProviders(query string) []SearchResult {
	var searchResults []SearchResult

	async.Range(e.providers, func(p Provider) {
		resultsForProvider := p.Search(query)
		searchResults = append(searchResults, resultsForProvider...)
	})

	return searchResults
}
