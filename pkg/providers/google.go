package providers

import (
	"fmt"
	"log"
	"searchengine/pkg/entities"
	"searchengine/pkg/net"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Google struct {
	HttpClient net.HttpClient
}

const GOOGLE_URL = "https://google.com/search"

func (p Google) Search(query string) []entities.SearchResult {
	var results []entities.SearchResult
	res, err := p.HttpClient.DoCall(fmt.Sprintf("%s?q=%s", GOOGLE_URL, query))

	if err != nil {
		log.Println(err)
	}

	document, _ := goquery.NewDocumentFromReader(res.Body)
	elements := document.Find("div.g")

	for index := range elements.Nodes {
		item := elements.Eq(index)

		results = append(results, googleParsing(item))
	}

	return results
}

func googleParsing(item *goquery.Selection) entities.SearchResult {
	link := strings.TrimSpace(item.Find("a").AttrOr("href", ""))
	title, _ := item.Find("h3").First().Html()
	description, _ := item.Find(".VwiC3b").Html()

	return entities.SearchResult{
		Title:       title,
		Description: description,
		Link:        link,
		From:        "Google",
	}
}
