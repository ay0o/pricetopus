package parser

import "strings"

type selector struct {
	name  string
	price string
}

var (
	selectors = map[string]selector{
		"amazon": {
			name:  "#productTitle",
			price: ".priceBlockBuyingPriceString"},
		"ebay": {
			name:  "title",
			price: "span#prcIsum"},
		"pccomponentes": {
			name:  "title",
			price: "#precio-main"},
	}
)

func getSelectors(host string) (string, string) {
	for key, value := range selectors {
		if strings.Contains(host, key) {
			return value.name, value.price
		}
	}
	return "", ""
}
