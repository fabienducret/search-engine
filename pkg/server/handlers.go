package server

import (
	"html/template"
	"net/http"
	"net/url"
	"path"
	"searchengine/pkg/domain"
	"searchengine/pkg/viewmodel"
)

func (s *Server) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	searchHandler := searchHandlerFactory(s.Engine)
	mux.HandleFunc("/", searchHandler)
	mux.HandleFunc("/search", searchHandler)

	return mux
}

func searchHandlerFactory(engine domain.Engine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles(path.Join("web", "templates", "search.html"))

		query := url.QueryEscape(r.URL.Query().Get("q"))
		var searchResults []domain.SearchResult

		if query != "" {
			searchResults = engine.Search(query)
		}

		response := viewmodel.Response(query, searchResults)

		t.Execute(w, response)
	}
}
