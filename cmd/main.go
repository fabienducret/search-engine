package main

import (
	"log"
	"net/http"
	"searchengine/pkg/domain"
	"searchengine/pkg/infra"
	"searchengine/pkg/server"
)

func main() {
	google := infra.Google{}
	engine := domain.NewEngine([]domain.Provider{google})

	log.Fatal(http.ListenAndServe(":3333", server.New(engine)))
}
