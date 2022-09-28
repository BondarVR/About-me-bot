package main

import (
	"about-me/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"

	"about-me/internal/config"
	"about-me/internal/telegram"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	lgr, err := logger.New(logger.Config{
		LogLevel:    cfg.LogLevel,
		LogServer:   cfg.LogServer,
		ServiceName: cfg.ServiceName,
	})
	if err != nil {
		lgr.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		lgr.Fatalf("tgbotapi.NewBotAPI() failed. Error: '%v'\n", err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot, lgr)
	if err := telegramBot.Start(); err != nil {
		lgr.Fatal(err)
	}
}
