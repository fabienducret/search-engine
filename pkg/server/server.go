package server

import (
	"html/template"
	"net/http"
	"path"
	"searchengine/pkg/domain"
	"searchengine/pkg/viewmodel"
)

func New(engine domain.Engine) *http.ServeMux {
	server := http.NewServeMux()

	searchHandler := searchHandlerFactory(engine)
	server.HandleFunc("/", searchHandler)
	server.HandleFunc("/search", searchHandler)

	return server
}

func searchHandlerFactory(engine domain.Engine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles(path.Join("web", "templates", "search.html"))

		query := r.URL.Query().Get("q")

		searchResults := engine.Search(query)
		response := viewmodel.Response(query, searchResults)

		t.Execute(w, response)
	}
}
