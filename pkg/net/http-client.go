package net

import (
	"net/http"
)

type HttpClient struct {
	http.Client
}

func (c HttpClient) DoCall(url string) (*http.Response, error) {
	req := request(url)
	return c.Do(req)
}

func request(url string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "fr,fr-FR;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36")

	return req
}
