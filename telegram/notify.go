package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hekmon/transmissionrpc/v3"
	"torrent-listener/infrastructure"
)

func Notify(torrent transmissionrpc.Torrent) {
	var messageText = fmt.Sprintf(infrastructure.Config().TelegramConfig.DownloadCompletedMessage, *torrent.Name)
	msg := tgbotapi.NewMessage(infrastructure.Config().TelegramConfig.ChatID, messageText)

	_, err := BotAPI().Send(msg)
	if err != nil {
		fmt.Println(err.Error())
	}
}
