package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

type PostalCode int

type Advertisement struct {
	Name       string
	URL        string
	Quarter    string
	PostalCode PostalCode
}

func main() {
	app := Application{
		InfoLogger:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.InfoLogger.Println("Starting to scrape")

	c := setupColly()

	url := "https://www.kleinanzeigen.de/s-wohnung-mieten/hamburg/c203l9409"
	err := c.Visit(url)
	if err != nil {
		app.ErrorLogger.Fatalf("Fetching the url %q returned error", url)
	}
}

func setupColly() *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	c.OnHTML(".aditem-main", func(element *colly.HTMLElement) {
		fmt.Println(element.Index)
	})

	return c
}
