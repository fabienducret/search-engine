package engine_test

import (
	"searchengine/pkg/engine"
	"searchengine/pkg/entities"
	"testing"
)

type FirstFakeProvider struct{}

func (p FirstFakeProvider) Search(q string) []entities.SearchResult {
	return []entities.SearchResult{
		{Title: "fake first title", Description: "fake first description", Link: "www.firstlink.com"},
	}
}

type SecondFakeProvider struct{}

func (p SecondFakeProvider) Search(q string) []entities.SearchResult {
	return []entities.SearchResult{
		{Title: "fake second title", Description: "fake second description", Link: "www.secondlink.com"},
		{Title: "fake third title", Description: "fake third description", Link: "www.thirstlink.com"},
	}
}

type ThirdFakeProvider struct{}

func (p ThirdFakeProvider) Search(q string) []entities.SearchResult {
	return []entities.SearchResult{
		{Title: "fake third title", Description: "fake third description", Link: "www.thirstlink.com"},
		{Title: "fake first title", Description: "fake first description", Link: "www.firstlink.com"},
	}
}

func TestEngine(t *testing.T) {
	t.Run("aggregates search results", func(t *testing.T) {
		// Given
		want := 3
		engine := engine.NewEngine([]engine.Provider{
			FirstFakeProvider{},
			SecondFakeProvider{},
		})

		// When
		got := len(engine.Search("query"))

		if want != got {
			t.Errorf("Invalid search response, want %d, got %d", want, got)
		}
	})

	t.Run("aggregates and deduplicates search results", func(t *testing.T) {
		// Given
		want := 3
		engine := engine.NewEngine([]engine.Provider{
			FirstFakeProvider{},
			SecondFakeProvider{},
			ThirdFakeProvider{},
		})

		// When
		got := len(engine.Search("query"))

		if want != got {
			t.Errorf("Invalid search response, want %d, got %d", want, got)
		}
	})
}
