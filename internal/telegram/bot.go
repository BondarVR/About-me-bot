package telegram

import (
	"about-me/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	lgr *logger.LogrusLogger
}

func NewBot(bot *tgbotapi.BotAPI, lgr *logger.LogrusLogger) *Bot {
	return &Bot{bot: bot,
		lgr: lgr,
	}
}

// Start (starts the bot)
func (b *Bot) Start() error {
	b.lgr.Infof("Authorized on account %s", b.bot.Self.UserName)
	updates := b.initUpdatesChannel()
	if err := b.handleUpdates(updates); err != nil {
		return err
	}
	return nil
}

// Processing updates
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				return err
			}
			continue
		}
		if err := b.handleMessage(update.Message); err != nil {
			return err
		}
	}
	return nil
}

// Getting updates
func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
