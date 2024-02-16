package commands

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"torrent-listener/infrastructure"
)

const DefaultWelcomeMessage = "Hello!"

type Start struct {
}

func (l Start) Handle(message *tgbotapi.Message, api *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(message.Chat.ID, welcomeMessage())
	msg.ParseMode = "Markdown"

	_, err := api.Send(msg)
	if err != nil {
		_ = errors.New(err.Error())
	}
}

func welcomeMessage() string {
	welcomeMessage := infrastructure.Config().TelegramConfig.WelcomeMessage
	if len(welcomeMessage) == 0 {
		welcomeMessage = DefaultWelcomeMessage
	}

	return welcomeMessage
}
