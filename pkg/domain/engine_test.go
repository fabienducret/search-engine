package domain_test

import (
	"searchengine/pkg/domain"
	"testing"
)

type FirstFakeProvider struct{}

func (p FirstFakeProvider) Search(q string) []domain.SearchResult {
	return []domain.SearchResult{
		{Title: "fake first title", Description: "fake first description", Link: "www.firstlink.com"},
	}
}

type SecondFakeProvider struct{}

func (p SecondFakeProvider) Search(q string) []domain.SearchResult {
	return []domain.SearchResult{
		{Title: "fake second title", Description: "fake second description", Link: "www.secondlink.com"},
		{Title: "fake third title", Description: "fake third description", Link: "www.firstlink.com"},
	}
}

func TestEngine(t *testing.T) {
	t.Run("aggregates and deduplicates search results", func(t *testing.T) {
		// Given
		want := 2
		firstProvider := FirstFakeProvider{}
		secondProvider := SecondFakeProvider{}
		engine := domain.NewEngine([]domain.Provider{firstProvider, secondProvider})

		// When
		got := len(engine.Search("query"))

		if want != got {
			t.Errorf("Invalid search response, want %d, got %d", want, got)
		}
	})
}
