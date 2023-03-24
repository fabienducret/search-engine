FROM golang:1.20.2-bullseye

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN go build -o /search-engine ./cmd/main.go

EXPOSE 8080


ENTRYPOINT ["/search-engine"]