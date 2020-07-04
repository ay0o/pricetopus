package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"github.com/ay0o/pricetopus/internal/parser"
)

type email struct {
	server   string
	port     string
	user     string
	from     string
	password string
	to       string
}

func (e email) send(msg []byte) {

	auth := smtp.PlainAuth("", e.user, e.password, e.server)
	endpoint := e.server + ":" + e.port
	err := smtp.SendMail(endpoint, auth, e.from, []string{e.to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := email{
		server:   os.Getenv("PRICETOPUS_EMAIL_SERVER"),
		port:     os.Getenv("PRICETOPUS_EMAIL_SERVER_PORT"),
		user:     os.Getenv("PRICETOPUS_EMAIL_USER"),
		from:     os.Getenv("PRICETOPUS_EMAIL_USER"),
		password: os.Getenv("PRICETOPUS_EMAIL_PASSWORD"),
		to:       os.Getenv("PRICETOPUS_EMAIL_TO"),
	}

	expectedPrice, err := strconv.ParseFloat(os.Getenv("PRICETOPUS_PRODUCT_PRICE"), 64)
	if err != nil {
		log.Fatal(err)
	}
	url := os.Getenv("PRICETOPUS_PRODUCT_URL")

	name, price := parser.Parse(url)

	if price <= expectedPrice {
		msg := []byte("Subject: [Pricetopus] " + name + "\n" + "Product: " + name + "\n\n" + "Price: " + fmt.Sprintf("%0.2f", price))
		e.send(msg)
	}

	log.Printf("Product: %s ---- Price: %0.2f\n", name, price)
}
