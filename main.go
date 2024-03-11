package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const adFileName = "advertisements.json" // TODO from env

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	AdMap       map[string]Advertisement
}

func main() {
	app := newApplication()

	app.initializeAdsFromFile(adFileName)
	app.scrape()
	app.writeAdMapToJson()

	fmt.Println("Successfully scraped everything")
}

func (app *Application) writeAdMapToJson() {
	adAsJsonByte, _ := json.Marshal(app.AdMap)

	err := os.WriteFile("advertisements.json", adAsJsonByte, 0666)
	if err != nil {
		app.ErrorLogger.Printf("Error when writing to file: %s", err)
	}
}

func newApplication() *Application {
	app := Application{
		InfoLogger:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
		AdMap:       make(map[string]Advertisement),
	}

	return &app
}

func (app *Application) addAd(ad Advertisement) {
	_, exists := app.AdMap[ad.URL]
	if exists {
		app.InfoLogger.Printf("Didnt add advertisement, URL existed already: %s\n", ad.URL)
	} else {
		ok := app.checkIfAdSuitable(ad)
		if ok {
			app.InfoLogger.Printf("Adding advertisement: %+v\n", ad)
			app.AdMap[ad.URL] = ad
		}
	}
}

func (app *Application) scrape() {
	app.InfoLogger.Println("Starting to scrape")
	app.initKleinanzeigenScraper()
}

func (app *Application) initializeAdsFromFile(name string) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		app.ErrorLogger.Printf("Error during loading ad map: %s", err)
	}
	err = json.Unmarshal(bytes, &app.AdMap)
}
