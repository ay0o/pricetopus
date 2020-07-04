package parser

import (
	"log"
	"net/url"
)

func Parse(site string) (string, float64) {
	u, err := url.Parse(site)
	if err != nil {
		log.Fatal(err)
	}

	nameSelector, priceSelector := getSelectors(u.Host)
	if nameSelector == "" {
		log.Fatal("Unsupported site")
	}

	doc := getDocument(site)
	name, err := getName(doc, nameSelector)
	if err != nil {
		log.Fatal(err)
	}
	price, err := getPrice(doc, priceSelector)
	if err != nil {
		log.Fatal(err)
	}
	return name, price
}
