package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

const kleinanzeigenURL = "https://www.kleinanzeigen.de"

func (app *Application) initKleinanzeigenScraper() {
	app.InfoLogger.Println("Starting to scrape kleinanzeigen...")

	c := setupColly()

	c.OnHTML(".aditem-main", func(element *colly.HTMLElement) {
		locationString := element.ChildText(".aditem-main--top--left")
		postalCode, quarter := app.getPostalCodeAndQuarter(locationString)
		adPath := element.ChildAttr(".ellipsis", "href")
		adName := element.ChildText(".ellipsis")
		ad := Advertisement{
			Name:       adName,
			URL:        kleinanzeigenURL + adPath,
			Quarter:    quarter,
			PostalCode: postalCode,
		}
		app.addAd(ad)
	})

	url := kleinanzeigenURL + "/s-wohnung-mieten/hamburg/c203l9409"
	err := c.Visit(url)
	if err != nil {
		app.ErrorLogger.Printf("Fetching the url %q returned error: %v", url, err)
	}
}

func (app *Application) getPostalCodeAndQuarter(locationString string) (int, string) {
	//Kleinanzeigen string is ie "20095 Hamburg Altstadt"
	parts := strings.Split(locationString, " ")
	var quarter string
	// ie if we only have "21073 Harburg"
	if len(parts) < 3 {
		quarter = "NO QUARTER"
	} else {
		quarter = parts[2]
	}
	postalCodeString := parts[0]

	//ie St. Georg
	if len(parts) == 4 {
		quarter += parts[3]
	}

	postalCode, err := strconv.Atoi(postalCodeString)
	if err != nil {
		logMessage := fmt.Sprintf("Conversion of postal code didnt work: %q, err: %v", postalCodeString, err)
		app.InfoLogger.Println(logMessage)
	}

	return postalCode, quarter
}
