package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	AdList      []Advertisement
}

func main() {
	app := newApplication()

	app.InfoLogger.Println("Starting to scrape")
	app.initKleinanzeigenScraper()

	writeAdListToJson(app)

	fmt.Println("Successfully scraped everything")
}

func writeAdListToJson(app *Application) {
	adAsJsonByte, _ := json.Marshal(app.AdList)

	err := os.WriteFile("advertisements.json", adAsJsonByte, 0666)
	if err != nil {
		app.ErrorLogger.Printf("Error when writing to file: %s", err)
	}
}

func newApplication() *Application {
	app := Application{
		InfoLogger:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
		AdList:      []Advertisement{},
	}

	return &app
}

func (app *Application) addAd(advertisement Advertisement) {
	suitableAd := checkAdvertisementCriteria(advertisement)
	if suitableAd {
		app.InfoLogger.Printf("Adding advertisement: %+v\n", advertisement)
		app.AdList = append(app.AdList, advertisement)
	} else {
		app.InfoLogger.Printf("Didnt add advertisement: %+v\n", advertisement)
	}
}
