package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"golang.org/x/exp/maps"
	"log"
	"os"
	"time"
)

const adFileName = "advertisements.json" // TODO from env

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	AdMap       map[string]Advertisement
	NewAds      map[string]Advertisement
	TelegramBot *bot.Bot
}

func main() {
	app := newApplication()
	app.initializeAdsFromFile(adFileName)

	ready := make(chan bool)
	go app.StartBot(getDotEnvVariable("TELEGRAM_API_KEY"), ready)
	<-ready

	ticker := time.NewTicker(3 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		app.scrape()
		app.writeAdMapToJson()
		app.sendLastAdsAsMessage()
		app.mergeAdLists()
	}

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
		NewAds:      make(map[string]Advertisement),
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
			app.NewAds[ad.URL] = ad
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

func getDotEnvVariable(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func (app *Application) sendLastAdsAsMessage() {
	//last count entries from json? cuz in map we cant know which are the last entries
	for _, ad := range app.NewAds {
		message := buildMessageFromAd(ad)
		app.sendMessage(message)
	}
}

func (app *Application) mergeAdLists() {
	maps.Copy(app.AdMap, app.NewAds)
	app.NewAds = make(map[string]Advertisement)
}

func buildMessageFromAd(ad Advertisement) string {
	return fmt.Sprintf("%s\n in %s, %d\n%s", ad.Name, ad.Quarter, ad.PostalCode, ad.URL)
}
