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
			tgbotapi.NewKeyboardButton(model.Gear),
			tgbotapi.NewKeyboardButton(model.Info),
			tgbotapi.NewKeyboardButton(model.Help),
		},
	)
}

func unitsKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
			tgbotapi.NewKeyboardButton(l.Language[lang]["°C, mps"]),
			tgbotapi.NewKeyboardButton(l.Language[lang]["°F, mph"]),
		},
	)
}

func settingsKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
			tgbotapi.NewKeyboardButton(model.GlobeWithMeridian),
			tgbotapi.NewKeyboardButton(model.TriangularRuler),
		},
	)
}

func langKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
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
