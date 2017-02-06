package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	c "github.com/lavrs/telegram-weather-bot/config"
	msg "github.com/lavrs/telegram-weather-bot/message"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	"log"
)

func main() {
	c.SetConfig()

	bot, err := tgbotapi.NewBotAPI(c.Cfg.TelegramToken)
	errors.CheckErrPanic(err)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60

	updates, err := bot.GetUpdatesChan(upd)
	errors.CheckErrPanic(err)

	for update := range updates {
		msg.Updates(bot, update)
	}
}
