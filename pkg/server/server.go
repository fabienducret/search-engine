package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"searchengine/pkg/domain"
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
		p := path.Join("web", "templates", "search.html")
		t, _ := template.ParseFiles(p)

		query := r.URL.Query().Get("q")

		fmt.Println(engine.Search(query))

		injectedParams := map[string]interface{}{"query": query}

		t.Execute(w, injectedParams)
	}
}
