package parser

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var amazonSampleCode = `
<html>
	<body>
		<span id="productTitle">Amazon Test</span>
		<span class="priceBlockBuyingPriceString">123</span>
	</body>
</html>`

var ebaySampleCode = `
<html>
	<head>
		<title>Ebay Test</title>
		<span id="prcIsum">22.22</span>
	</head>
</html>`

var pccSampleCode = `
<html>
	<head>
		<title>PCComponentes Test</title>
		<div id="precio-main">0.75</div>
	</head>
</html>`

type siteTest struct {
	site   string
	source string
	name   string
	price  float64
}

var siteTests = []siteTest{
	{"amazon", amazonSampleCode, "Amazon Test", 123},
	{"ebay", ebaySampleCode, "Ebay Test", 22.22},
	{"pccomponentes", pccSampleCode, "PCComponentes Test", 0.75},
}

func TestGetName(t *testing.T) {
	for _, test := range siteTests {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(test.source))
		if err != nil {
			t.Fatalf("%s", err)
		}
		nameSelector, _ := getSelectors(test.site)
		name, err := getName(doc, nameSelector)
		if err != nil {
			t.Fatal(err)
		}
		if name != test.name {
			t.Errorf("[%s] Expected: %s. Got: %s", test.site, test.name, name)
		}
	}
}

func TestGetPrice(t *testing.T) {
	for _, test := range siteTests {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(test.source))
		if err != nil {
			t.Fatalf("%s", err)
		}
		_, priceSelector := getSelectors(test.site)
		price, err := getPrice(doc, priceSelector)
		if err != nil {
			t.Fatal(err)
		}
		if price != test.price {
			t.Errorf("[%s] Expected: %f. Got: %f", test.site, test.price, price)
		}
	}
}
