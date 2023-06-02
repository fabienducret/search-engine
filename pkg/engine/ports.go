package engine

import "searchengine/pkg/entities"

type Provider interface {
	Search(q string) []entities.SearchResult
}
