package server

import (
	"html/template"
	"net/http"
	"path"
)

func New() *http.ServeMux {
	server := http.NewServeMux()

	server.HandleFunc("/", searchHandler)
	server.HandleFunc("/search", searchHandler)

	return server
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	p := path.Join("web", "templates", "search.html")
	t, _ := template.ParseFiles(p)

	query := r.URL.Query().Get("q")
	injectedParams := map[string]interface{}{"query": query}

	t.Execute(w, injectedParams)
}
