package server

import "searchengine/pkg/entities"

type Engine interface {
	Search(query string) []entities.SearchResult
}
