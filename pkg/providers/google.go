package providers

import (
	"fmt"
	"searchengine/pkg/domain"
	"searchengine/pkg/net"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Google struct {
	HttpClient net.HttpClient
}

func (p Google) Search(query string) []domain.SearchResult {
	var results []domain.SearchResult
	res, _ := p.HttpClient.DoCall(fmt.Sprintf("http://google.com/search?q=%s", query))

	document, _ := goquery.NewDocumentFromReader(res.Body)
	elements := document.Find("div.g")

	for index := range elements.Nodes {
		item := elements.Eq(index)

		results = append(results, domain.SearchResult{
			Title:       titleFrom(item),
			Description: descriptionFrom(item),
			Link:        linkFrom(item),
			From:        "Google",
		})
	}

	return results
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
