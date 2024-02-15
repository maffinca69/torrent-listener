package commands

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math"
	"sort"
	"strconv"
	"torrent-listener/transmission"
)

const TotalCountForShow = 5

type List struct {
}

func (l List) Handle(message *tgbotapi.Message, api *tgbotapi.BotAPI) {
	var messageText = getTorrents()
	msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = message.MessageID

	_, err := api.Send(msg)
	if err != nil {
		println(err.Error())
	}
}

func getTorrents() string {
	torrents, _ := transmission.Client().TorrentGetAll(context.TODO())
	if len(torrents) == 0 {
		return "Not found torrents"
	}

	sort.Slice(torrents, func(i, j int) bool {
		return torrents[i].AddedDate.Unix() > torrents[j].AddedDate.Unix()
	})

	// Get latest 5 torrents
	splitTorrents := torrents[0:TotalCountForShow]

	var text = ""
	for _, torrent := range splitTorrents {
		progress := math.RoundToEven(*torrent.PercentDone * 100)

		text += "💿 *" + *torrent.Name + "*" + "\n"
		text += "Hash: " + *torrent.HashString + "\n"
		text += "Size: " + torrent.TotalSize.String() + "\n"
		text += "Status: " + torrent.Status.String() + "\n"
		text += "Progress: " + strconv.Itoa(int(progress)) + "%"

		text += "\n\n"
	}

	return text
}
