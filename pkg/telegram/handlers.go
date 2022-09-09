package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	commandHelp  = "help"
	commandAbout = "about"
	commandLinks = "links"

	replyStartAndHelp = "Here is a list of commands that I can handle\n" +
		"/start - show a list of commands\n" +
		"/help - show a list of commands\n" +
		"/about - brief information about the owner\n" +
		"/links - list of social media links\n"

	replyAbout = "My name is Vlad, I'm 20 years old.\n" +
		"I am a 4th year student at the Polytechnic University.\n" +
		"I am learning the golang programming language.\n" +
		"In the future I want to become a software engineer\n\n" +
		"To see the available commands, send /help."

	replyLinks = "My social networks:\n" +
		"Linkedin - http://www.linkedin.com/in/vladyslav-bondar-68192b24b\n" +
		"Github - https://github.com/BondarVR\n" +
		"Codewars - https://www.codewars.com/users/Bondar_VR/badges/micro\n" +
		"Facebook - https://www.facebook.com/profile.php?id=100008469645353\n" +
		"Telegram - https://t.me/VladislavBondar\n\n" +
		"To see the available commands, send /help."
)

// Process messages from the update and send back
func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't understand what to do, try enter /help")
	_, err := b.bot.Send(msg)
	return err
}

// Process command from the update and send back
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "There is no such command, enter /help")
	switch message.Command() {
	case commandStart:
		msg.Text = replyStartAndHelp
		//msg.ReplyToMessageID = message.MessageID
		_, err := b.bot.Send(msg)
		return err
	case commandHelp:
		msg.Text = replyStartAndHelp
		msg.ReplyToMessageID = message.MessageID
		_, err := b.bot.Send(msg)
		return err
	case commandAbout:
		msg.Text = replyAbout
		_, err := b.bot.Send(msg)
		return err
	case commandLinks:
		msg.Text = replyLinks
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}
