package domain

import "sync"

type Engine struct {
	providers []Provider
}

func NewEngine(providers []Provider) Engine {
	return Engine{providers: providers}
}

func (e Engine) Search(query string) []SearchResult {
	return e.searchInParallel(query)
}

func (e Engine) searchInParallel(query string) []SearchResult {
	var searchResults []SearchResult
	wg := sync.WaitGroup{}

	for _, p := range e.providers {
		wg.Add(1)

		go func(p Provider) {
			defer wg.Done()

			results := p.Search(query)
			searchResults = append(searchResults, results...)
		}(p)
	}
	wg.Wait()

	return searchResults
}
