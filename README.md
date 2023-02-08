# search-engine

## Introduction

This project is an aggregator of several search engines.

The purpose is not to build something unique or valuable but to learn.

## How to use

Run the command :

`go run cmd/main.go`

Then browse this URL : http://localhost:3333/

## Structure

```
project
│   README.md
|   go.mod
|   go.sum
│
└───cmd
│   │   main.go
│
└───pkg
│   └───async
│   |   │   range.go
│   |
│   └───domain
│   |   │   engine_test.go
│   |   │   engine.go
│   |   │   provider.go
│   |   │   searchresult.go
│   |   │   range.go
│   |
│   └───net
│   |   │   http-client.go
│   |
│   └───providers
│   |   │   bing.go
│   |   │   google.go
│   |
│   └───server
│   |   │   server.go
│   |
│   └───viewmodel
│   |   │   response.go
│
└───web
│   └───templates
│   |   │   search.html
```

## Tests

Run the command :

`go test ./... `

## Roadmap

This project is under construct. Here is the roadmap :

- [x] Add come from provider in search results
- [ ] Add Bing as provider
- [ ] Add DuckDuckGo as provider
