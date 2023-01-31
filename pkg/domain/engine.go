package domain

type Engine struct {
	providers []Provider
}

func NewEngine(providers []Provider) Engine {
	return Engine{providers: providers}
}

func (e Engine) Search(query string) []SearchResult {
	var searchResults []SearchResult
	providers := e.providers

	// should be in parallel
	for _, p := range providers {
		resultsForProvider := p.Search(query)
		searchResults = append(searchResults, resultsForProvider...)
	}

	return searchResults
}
