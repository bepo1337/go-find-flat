package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"os"
	"os/signal"
)

const channelName = "@GoFindFlatBot"

func (app *Application) StartBot(apikey string, ready chan<- bool) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	b, _ := bot.New(apikey)
	app.TelegramBot = b
	ready <- true
	close(ready)
	b.Start(ctx)
}

func (app *Application) sendMesssage() {
	_, err := app.TelegramBot.SendMessage(context.TODO(), &bot.SendMessageParams{
		ChatID: "5629048775",
		Text:   "Hello from Go",
	})

	if err != nil {
		log.Printf("Error when sending telegram message %s", err)
	}
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	fmt.Println(update.Message.Chat.ID)
}
