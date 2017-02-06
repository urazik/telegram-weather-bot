package msg

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	l "github.com/lavrs/telegram-weather-bot/language"
	"github.com/lavrs/telegram-weather-bot/model"
)

func mainKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(l.Language[lang]["now"]),
			tgbotapi.NewKeyboardButton(l.Language[lang]["forToday"]),
		},
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(l.Language[lang]["forTomorrow"]),
			tgbotapi.NewKeyboardButton(l.Language[lang]["forWeek"]),
		},
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.GlobeWithMeridian),
			tgbotapi.NewKeyboardButton(model.Info),
			tgbotapi.NewKeyboardButton(model.Help),
		},
	)
}

func langKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.CountriesFATE["en"]),
			tgbotapi.NewKeyboardButton(model.CountriesFATE["ru"]),
		},
	)
}

func helpKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Help),
		},
	)
}
