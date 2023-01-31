package domain_test

import (
	"reflect"
	"searchengine/pkg/domain"
	"testing"
)

type FirstFakeProvider struct{}

func (p FirstFakeProvider) Search(q string) []domain.SearchResult {
	return []domain.SearchResult{
		{Title: "fake first title", Description: "fake first description", Link: "fake first link"},
	}
}

type SecondFakeProvider struct{}

func (p SecondFakeProvider) Search(q string) []domain.SearchResult {
	return []domain.SearchResult{
		{Title: "fake second title", Description: "fake second description", Link: "fake second link"},
	}
}

func TestEngine(t *testing.T) {
	t.Run("executes a provider", func(t *testing.T) {
		// Given
		want := []domain.SearchResult{
			{Title: "fake first title", Description: "fake first description", Link: "fake first link"},
			{Title: "fake second title", Description: "fake second description", Link: "fake second link"},
		}
		firstProvider := FirstFakeProvider{}
		secondProvider := SecondFakeProvider{}
		engine := domain.NewEngine([]domain.Provider{firstProvider, secondProvider})

		// When
		got := engine.Search("query")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Invalid search response, want %s, got %s", want, got)
		}
	})
}
