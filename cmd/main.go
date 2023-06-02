package main

import (
	"searchengine/pkg/engine"
	"searchengine/pkg/net"
	"searchengine/pkg/providers"
	"searchengine/pkg/server"
)

func main() {
	server.New(buildEngine()).Start()
}

func buildEngine() engine.Engine {
	httpClient := net.HttpClient{}
	return engine.NewEngine([]engine.Provider{
		providers.Google{
			HttpClient: httpClient,
		},
		providers.Bing{
			HttpClient: httpClient,
		},
	})
}
