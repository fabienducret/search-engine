package main

import (
	"log"
	"net/http"
	"searchengine/pkg/domain"
	"searchengine/pkg/providers"
	"searchengine/pkg/server"
)

func main() {
	google := providers.Google{}
	bing := providers.Bing{}
	engine := domain.NewEngine([]domain.Provider{google, bing})

	log.Fatal(http.ListenAndServe(":3333", server.New(engine)))
}
