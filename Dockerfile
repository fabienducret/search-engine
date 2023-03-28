FROM golang:alpine 

WORKDIR /app
COPY . ./

RUN go build -o /search-engine ./cmd/main.go

EXPOSE 8080

ENTRYPOINT ["/search-engine"]