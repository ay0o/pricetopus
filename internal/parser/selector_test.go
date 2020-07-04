package parser

import "testing"

func TestGetSelectors(t *testing.T) {
	name, price := getSelectors("www.amazon.com")
	if name != "#productTitle" && price != ".priceBlockBuyingPriceString" {
		t.Error("Problem getting selectors")
	}
}
