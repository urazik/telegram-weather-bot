package errors

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/telegram-weather-bot/config"
	"log"
)

func CheckErrPanic(cerr error) {
	if cerr != nil {
		bot, _ := tgbotapi.NewBotAPI(config.Cfg.TelegramToken)
		msg := tgbotapi.NewMessage(config.Cfg.MyTelegramID, cerr.Error())
		bot.Send(msg)
		log.Panic(cerr)
	}
}
