package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"torrent-listener/telegram/commands"
)

type Command interface {
	Handle(message *tgbotapi.Message, api *tgbotapi.BotAPI)
}

func ResolveCommand(command string) Command {
	switch command {
	case "list":
		return commands.List{}
	case "ping":
		return commands.Ping{}
	default:
		fmt.Println("No implemented yet: " + command)
		return nil
	}
}
