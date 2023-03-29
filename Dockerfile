FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /app/search-engine ./cmd

FROM scratch AS final

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app ./

EXPOSE 8080

ENTRYPOINT ["/app/search-engine"]