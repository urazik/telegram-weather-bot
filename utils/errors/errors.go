package errors

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/telegram-weather-bot/config"
	"log"
)

func CheckErrPanic(cerr error) {
	if cerr != nil {
		bot, err := tgbotapi.NewBotAPI(config.Cfg.TelegramToken)
		if err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(config.Cfg.MyTelegramID, cerr.Error())
		_, err = bot.Send(msg)
		if err != nil {
			log.Println(err)
		}

		log.Panic(cerr)
	}
}
