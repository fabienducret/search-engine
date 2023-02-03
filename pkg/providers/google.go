package providers

import (
	"fmt"
	"net/http"
	"searchengine/pkg/domain"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Google struct {
}

func (p Google) Search(query string) []domain.SearchResult {
	var results []domain.SearchResult
	req := request(query)
	client := &http.Client{}
	res, _ := client.Do(req)

	document, _ := goquery.NewDocumentFromReader(res.Body)
	elements := document.Find("div.g")

	for index := range elements.Nodes {
		item := elements.Eq(index)

		results = append(results, domain.SearchResult{
			Title:       titleFrom(item),
			Description: descriptionFrom(item),
			Link:        linkFrom(item),
		})
	}

	return results
}

func request(query string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://google.com/search?q=%s", query), nil)

	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "fr,fr-FR;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36")

	return req
}

func linkFrom(item *goquery.Selection) string {
	a := item.Find("a")
	return strings.TrimSpace(a.AttrOr("href", ""))
}

func titleFrom(item *goquery.Selection) string {
	selector := item.Find("h3").First()
	title, _ := selector.Html()

	return title
}

func descriptionFrom(item *goquery.Selection) string {
	selector := item.Find(".VwiC3b")
	description, _ := selector.Html()

	return description
}
