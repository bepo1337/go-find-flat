package main

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
	"os"
	"os/signal"
)

func (app *Application) StartBot(apikey string, ready chan<- bool) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	b, _ := bot.New(apikey)
	app.TelegramBot = b
	ready <- true
	close(ready)
	b.Start(ctx)
}

func (app *Application) sendMessage(message string) {
	_, err := app.TelegramBot.SendMessage(context.TODO(), &bot.SendMessageParams{
		ChatID: getChatID(),
		Text:   message,
	})

	if err != nil {
		log.Printf("Error when sending telegram message %s", err)
	}
}

func getChatID() string {
	return getDotEnvVariable("CHAT_ID")
}
