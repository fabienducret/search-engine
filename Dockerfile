FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /app/search-engine ./cmd

FROM gcr.io/distroless/static AS final

WORKDIR /app
COPY --from=builder /app ./

EXPOSE 8080

ENTRYPOINT ["/app/search-engine"]