package errors

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/telegram-weather-bot/config"
	"log"
)

func Check(err error) {
	if err != nil {
		bot, _ := tgbotapi.NewBotAPI(config.Cfg.TelegramToken)
		msg := tgbotapi.NewMessage(config.Cfg.MyTelegramID, err.Error())
		bot.Send(msg)
		log.Panic(err)
	}
}
