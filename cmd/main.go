package main

import (
	"log"
	"net/http"
	"searchengine/pkg/server"
)

func main() {
	log.Fatal(http.ListenAndServe(":3333", server.New()))
}
