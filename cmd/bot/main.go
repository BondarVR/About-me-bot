package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"about-me/pkg/config"
	"about-me/pkg/telegram"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cfg.Token: '%v'\n", cfg.Token)

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
