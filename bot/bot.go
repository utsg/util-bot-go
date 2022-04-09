package bot

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	util "github.com/utsg/util-bot-go/util"
)

func executeCommand(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "ip":
		msg.Text = util.GetIp()
	}

	return msg
}

func addTorrent(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	util.AddTorrent(update.Message.Text[1:len(update.Message.Text)-1])
	return msg
}

func RunBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil && util.IsUserAllowed(update.Message.From.UserName) {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				bot.Send(executeCommand(update))
			}

			if strings.HasPrefix(update.Message.Text, "+")  {
				bot.Send(addTorrent(update))
			}
		}
	}
}
