package msg

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/telegram-weather-bot/db"
	l "github.com/lavrs/telegram-weather-bot/language"
	"github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	"github.com/lavrs/telegram-weather-bot/utils/geocode"
	w "github.com/lavrs/telegram-weather-bot/weather"
)

func InfoMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	var msg tgbotapi.MessageConfig

	if user.Location == "" {
		msg = tgbotapi.NewMessage(telegramID,
			"*"+l.Language[user.Lang]["YourLL"]+"*\n"+
				"`"+l.Language[user.Lang]["empty_location"]+"`   "+model.CountriesFATE[user.Lang])
	} else {
		msg = tgbotapi.NewMessage(telegramID,
			"*"+l.Language[user.Lang]["YourLL"]+"*\n"+
				"`"+user.Location+"`   "+model.CountriesFATE[user.Lang])
	}

	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.CheckErrPanic(err)
}

func UpdateLangMsg(bot *tgbotapi.BotAPI, telegramID int64, message string) {
	isAuth, user := db.IsAuth(telegramID)

	var msg tgbotapi.MessageConfig

	if isAuth {
		lang := db.UpdateUserLang(user, model.CountriesFETA[message], telegramID)

		msg = tgbotapi.NewMessage(telegramID, l.Language[lang]["changeLanguageTo"]+" "+model.CountriesFATE[model.CountriesFETA[message]])
		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
	} else {
		db.SetUser(telegramID, nil, model.CountriesFETA[message])

		msg = tgbotapi.NewMessage(telegramID, l.Language[model.CountriesFETA[message]]["changeLanguageTo"]+" "+model.CountriesFATE[model.CountriesFETA[message]])
		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
	}

	_, err := bot.Send(msg)
	errors.CheckErrPanic(err)
}

func LangKeyboardMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)

	var msg tgbotapi.MessageConfig

	if isAuth {
		msg = tgbotapi.NewMessage(telegramID, l.Language[user.Lang]["chooseLanguage"])
	} else {
		msg = tgbotapi.NewMessage(telegramID, l.ChooseLanguage)
	}

	msg.ReplyMarkup = langKeyboard()
	_, err := bot.Send(msg)
	errors.CheckErrPanic(err)
}

func StartMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, _ := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	} else {
		Help(bot, telegramID)
	}
}

func Help(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	msg := tgbotapi.NewMessage(telegramID, l.Language[user.Lang]["help"])
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.CheckErrPanic(err)
}

func WeatherMsgFromCity(bot *tgbotapi.BotAPI, telegramID int64, location string) {
	isAuth, user := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	var msg tgbotapi.MessageConfig

	g, err := geocode.GetGeocode(location, user.Lang)
	if err != nil {
		msg = tgbotapi.NewMessage(telegramID, err.Error())
	} else {
		if user.Location != g.Result[0].FormattedAddress {
			db.SetUser(telegramID, g, user.Lang)

			msg := tgbotapi.NewMessage(telegramID, l.Language[user.Lang]["changeCityTo"]+" "+g.Result[0].FormattedAddress)
			msg.ReplyMarkup = mainKeyboard(user.Lang)
			_, err = bot.Send(msg)
			errors.CheckErrPanic(err)
		}

		wthr := w.CurrentWeather(
			g.Result[0].Geometry.Location.Lat, g.Result[0].Geometry.Location.Lng,
			user.Lang, g.Result[0].FormattedAddress)

		msg = tgbotapi.NewMessage(telegramID, wthr)
	}

	msg.ParseMode = "markdown"
	_, err = bot.Send(msg)
	errors.CheckErrPanic(err)
}

func WeatherMsgFromLocation(bot *tgbotapi.BotAPI, telegramID int64, location *tgbotapi.Location) {
	isAuth, user := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	var msg tgbotapi.MessageConfig

	g, err := geocode.GetReverseGeocode(location, user.Lang)
	if err != nil {
		msg = tgbotapi.NewMessage(telegramID, err.Error())
	} else {
		if (user.Lat != g.Result[0].Geometry.Location.Lat) || (user.Lng != g.Result[0].Geometry.Location.Lng) {
			db.SetUser(telegramID, g, user.Lang)

			msg = tgbotapi.NewMessage(telegramID, l.Language[user.Lang]["changeCityTo"]+" "+g.Result[0].FormattedAddress)
			msg.ReplyMarkup = mainKeyboard(user.Lang)
			_, err = bot.Send(msg)
			errors.CheckErrPanic(err)
		}

		wthr := w.CurrentWeatherFromLocation(user.Lang, location, g.Result[0].FormattedAddress)
		msg = tgbotapi.NewMessage(telegramID, wthr)
	}

	msg.ParseMode = "markdown"
	_, err = bot.Send(msg)
	errors.CheckErrPanic(err)
}

func WeatherMsgFromCmd(bot *tgbotapi.BotAPI, telegramID int64, message string) {
	isAuth, user := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	var (
		msg  tgbotapi.MessageConfig
		wthr string
		err  error
	)

	if user.Location == "" {
		msg = tgbotapi.NewMessage(telegramID, l.Language[user.Lang]["emptycity"])
		msg.ReplyMarkup = helpKeyboard()
	} else {
		switch {
		case (message == "now") || (message == "/now") || (message == "сейчас"):
			wthr = w.CurrentWeather(user.Lat, user.Lng, user.Lang, user.Location)

		case (message == "for today") || (message == "/today") || (message == "на сегодня"):
			wthr, err = w.WeatherOfDay(user)
			errors.CheckErrPanic(err)

		case (message == "for tomorrow") || (message == "/tomorrow") || (message == "на завтра"):
			wthr, err = w.TomorrowWeather(user)
			errors.CheckErrPanic(err)

		case (message == "for week") || (message == "/week") || (message == "на неделю"):
			wthr, err = w.WeekWeather(user)
			errors.CheckErrPanic(err)
		}

		msg = tgbotapi.NewMessage(telegramID, wthr)
		msg.ReplyMarkup = mainKeyboard(user.Lang)
		msg.ParseMode = "markdown"
	}

	_, err = bot.Send(msg)
	errors.CheckErrPanic(err)
}
