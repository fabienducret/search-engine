package providers

import (
	"fmt"
	"log"
	"searchengine/pkg/entities"
	"searchengine/pkg/net"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Bing struct {
	HttpClient net.HttpClient
}

const BING_URL = "https://www.bing.com/search"

func (p Bing) Search(query string) []entities.SearchResult {
	var results []entities.SearchResult
	res, err := p.HttpClient.DoCall(fmt.Sprintf("%s?q=%s", BING_URL, query))

	if err != nil {
		log.Println(err)
	}

	document, _ := goquery.NewDocumentFromReader(res.Body)
	elements := document.Find("li.b_algo")

	for index := range elements.Nodes {
		item := elements.Eq(index)

		results = append(results, bingParsing(item))
	}

	return results
}

func bingParsing(item *goquery.Selection) entities.SearchResult {
	link := strings.TrimSpace(item.Find("a").AttrOr("href", ""))
	title, _ := item.Find("h2").Find("a").Html()
	description, _ := item.Find("p.b_lineclamp2").Html()

	return entities.SearchResult{
		Title:       title,
		Description: description,
		Link:        link,
		From:        "Bing",
	}
}
