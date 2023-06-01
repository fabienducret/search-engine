package main

import (
	"searchengine/pkg/domain"
	"searchengine/pkg/net"
	"searchengine/pkg/providers"
	"searchengine/pkg/server"
)

func main() {
	server.New(engine()).Start()
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
