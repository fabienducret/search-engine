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

	return removeDuplicateLinks(searchResults)
}

func (e Engine) resultsFromProviders(query string) []SearchResult {
	var searchResults []SearchResult

	async.Range(e.providers, func(p Provider) {
		resultsForProvider := p.Search(query)
		searchResults = append(searchResults, resultsForProvider...)
	})

	return searchResults
}

func removeDuplicateLinks(toDeduplicate []SearchResult) []SearchResult {
	var searchResults []SearchResult

	for _, s := range toDeduplicate {
		if !contains(searchResults, s.Link) {
			searchResults = append(searchResults, s)
		}
	}

	return searchResults
}

func contains(searchResults []SearchResult, link string) bool {
	for _, s := range searchResults {
		if s.Link == link {
			return true
		}
	}

	return false
}
