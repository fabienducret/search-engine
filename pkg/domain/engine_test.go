package domain_test

import (
	"searchengine/pkg/domain"
	"testing"
)

type FakeProvider struct{}

func (p FakeProvider) Search(q string) string {
	return "fake"
}

func TestEngine(t *testing.T) {
	t.Run("executes a provider", func(t *testing.T) {
		// Given
		want := "fake"
		provider := FakeProvider{}
		engine := domain.NewEngine([]domain.Provider{provider})

		// When
		got := engine.Search("query")

		if want != got {
			t.Errorf("Invalid search response, want %s, got %s", want, got)
		}
	})
}
