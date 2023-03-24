package main

import (
	"log"
	"net/http"
	"searchengine/pkg/domain"
	"searchengine/pkg/net"
	"searchengine/pkg/providers"
	"searchengine/pkg/server"
)

func main() {
	engine := engine()
	log.Fatal(http.ListenAndServe(":8080", server.New(engine)))
}

func engine() domain.Engine {
	httpClient := net.HttpClient{}
	engine := domain.NewEngine([]domain.Provider{
		providers.Google{
			HttpClient: httpClient,
		},
		providers.Bing{
			HttpClient: httpClient,
		},
	})

	return engine
}
