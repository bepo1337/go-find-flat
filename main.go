package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

const adFileName = "advertisements.json" // TODO from env

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	AdMap       map[string]Advertisement
	TelegramBot *bot.Bot
}

func main() {
	app := newApplication()
	ready := make(chan bool)
	go app.StartBot(getDotEnvVariable("TELEGRAM_API_KEY"), ready)
	<-ready

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		app.sendMesssage()
		//TODO: every 5 min: scrape, check if new ads, if new --> send message
	}
	app.initializeAdsFromFile(adFileName)
	app.scrape()
	app.writeAdMapToJson()

	fmt.Println("Successfully scraped everything")
	if amIServer() {
		app.startServer()
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

func (app *Application) startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		adAsJsonByte, _ := json.Marshal(app.AdMap)
		_, err := w.Write(adAsJsonByte)
		if err != nil {
			app.ErrorLogger.Printf("Error when writing json to writer: %s", err)
		}
		return
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		app.ErrorLogger.Fatal("ListenAndServe: ", err)
	}
}

func amIServer() bool {
	if runtime.GOOS != "windows" {
		return true
	}
	return false
}

func getDotEnvVariable(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
