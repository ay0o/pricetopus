package parser

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	ua     string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"
	accept string = "text/html,application/xhtml+xml,application/xml"
)

func getDocument(url string) *goquery.Document {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("user-agent", ua)
	req.Header.Add("accept", accept)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func getName(doc *goquery.Document, selector string) (string, error) {
	name := doc.Find(selector).Text()
	if name == "" {
		return "", errors.New("Could not retrieve product name")
	}
	return strings.TrimSpace(name), nil
}

func getPrice(doc *goquery.Document, selector string) (float64, error) {
	priceStr := doc.Find(selector).Text()
	if priceStr == "" {
		return 0, errors.New("Could not retrieve product price")
	}

	priceStr = strings.ReplaceAll(priceStr, ",", ".")
	rgx := regexp.MustCompile(`([0-9]+[\\.]*[0-9]{1,2})`)

	currentPrice, err := strconv.ParseFloat(rgx.FindString(priceStr), 64)
	if err != nil {
		return 0, err
	}
	return currentPrice, nil
}
