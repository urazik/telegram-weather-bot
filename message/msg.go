package msg

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/google-geocode/geocode"
	c "github.com/lavrs/telegram-weather-bot/config"
	"github.com/lavrs/telegram-weather-bot/db"
	l "github.com/lavrs/telegram-weather-bot/language"
	"github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	w "github.com/lavrs/telegram-weather-bot/weather"
)

func InfoMsg(bot *tgbotapi.BotAPI, chatID int64) {
	isAuth, user := db.IsAuth(chatID)

	if !isAuth {
		LanguageMsg(bot, chatID)
		return
	}

	var msg tgbotapi.MessageConfig

	if user.Location == "" {
		msg = tgbotapi.NewMessage(chatID,
			"*"+l.Language[user.Lang]["YourLL"]+"*\n"+
				"`"+l.Language[user.Lang]["empty_location"]+"`   "+model.CountriesFATE[user.Lang])
	} else {
		msg = tgbotapi.NewMessage(chatID,
			"*"+l.Language[user.Lang]["YourLL"]+"*\n"+
				"`"+user.Location+"`   "+model.CountriesFATE[user.Lang])
	}

	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.CheckErrPanic(err)
}

func UpdateLangMsg(bot *tgbotapi.BotAPI, chatID int64, message string) {
	isAuth, user := db.IsAuth(chatID)

	var msg tgbotapi.MessageConfig

	if isAuth {
		lang := db.UpdateUserLang(user, model.CountriesFETA[message], chatID)

		msg = tgbotapi.NewMessage(chatID, l.Language[lang]["changeLanguageTo"]+" "+model.CountriesFATE[model.CountriesFETA[message]])
		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
	} else {
		db.SetUser(chatID, nil, model.CountriesFETA[message])

		msg = tgbotapi.NewMessage(chatID, l.Language[model.CountriesFETA[message]]["changeLanguageTo"]+" "+model.CountriesFATE[model.CountriesFETA[message]])
		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
	}

	_, err := bot.Send(msg)
	errors.CheckErrPanic(err)
}

func LanguageMsg(bot *tgbotapi.BotAPI, chatID int64) {
	isAuth, user := db.IsAuth(chatID)

	var msg tgbotapi.MessageConfig

	if isAuth {
		msg = tgbotapi.NewMessage(chatID, l.Language[user.Lang]["chooseLanguage"])
	} else {
		msg = tgbotapi.NewMessage(chatID, l.ChooseLanguage)
	}

	msg.ReplyMarkup = langKeyboard()
	_, err := bot.Send(msg)
	errors.CheckErrPanic(err)
}

func StartMsg(bot *tgbotapi.BotAPI, chatID int64) {
	isAuth, _ := db.IsAuth(chatID)

	if !isAuth {
		LanguageMsg(bot, chatID)
		return
	} else {
		Help(bot, chatID)
	}
}

func Help(bot *tgbotapi.BotAPI, chatID int64) {
	isAuth, user := db.IsAuth(chatID)

	if !isAuth {
		LanguageMsg(bot, chatID)
		return
	}

	msg := tgbotapi.NewMessage(chatID, l.Language[user.Lang]["help"])
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.CheckErrPanic(err)
}

func WeatherMsgFromCity(bot *tgbotapi.BotAPI, chatID int64, city string) {
	isAuth, user := db.IsAuth(chatID)

	if !isAuth {
		LanguageMsg(bot, chatID)
		return
	}

	var (
		g   *geocode.Geocoding
		err error
	)

	if user.Lang == "ru" {
		g, err = geocode.Geocode(city, geocode.Russian, c.Cfg.GoogleGeocodeToken)
	} else {
		g, err = geocode.Geocode(city, geocode.English, c.Cfg.GoogleGeocodeToken)
	}
	if err != nil {
		if err.Error() == "ZERO_RESULTS" {
			msg := tgbotapi.NewMessage(chatID, l.Language[user.Lang]["ZERO_RESULTS_CITY"])
			msg.ReplyMarkup = mainKeyboard(user.Lang)
			_, err := bot.Send(msg)
			errors.CheckErrPanic(err)

			return
		} else if err.Error() == "INVALID_REQUEST" {
			msg := tgbotapi.NewMessage(chatID, l.Language[user.Lang]["INVALID_REQUEST"])
			msg.ReplyMarkup = mainKeyboard(user.Lang)
			_, err := bot.Send(msg)
			errors.CheckErrPanic(err)

			return
		}

		errors.CheckErrPanic(err)
	}

	if user.Location != g.Result[0].FormattedAddress {
		db.SetUser(chatID, g, user.Lang)

		msg := tgbotapi.NewMessage(chatID, l.Language[user.Lang]["changeCityTo"]+" "+g.Result[0].FormattedAddress)
		msg.ReplyMarkup = mainKeyboard(user.Lang)
		_, err = bot.Send(msg)

		errors.CheckErrPanic(err)
	}

	wthr, err := w.CurrentWeather(
		g.Result[0].Geometry.Location.Lat, g.Result[0].Geometry.Location.Lng,
		user.Lang, g.Result[0].FormattedAddress)
	errors.CheckErrPanic(err)

	msg := tgbotapi.NewMessage(chatID, wthr)
	msg.ParseMode = "markdown"
	_, err = bot.Send(msg)
	errors.CheckErrPanic(err)
}

func WeatherMsgFromLocation(bot *tgbotapi.BotAPI, chatID int64, location *tgbotapi.Location) {
	isAuth, user := db.IsAuth(chatID)

	if !isAuth {
		LanguageMsg(bot, chatID)
		return
	}

	wthr, err := w.CurrentWeatherFromLocation(user.Lang, location)
	errors.CheckErrPanic(err)

	msg := tgbotapi.NewMessage(chatID, wthr)
	msg.ParseMode = "markdown"
	_, err = bot.Send(msg)
	errors.CheckErrPanic(err)
}

func WeatherMsgFromCmd(bot *tgbotapi.BotAPI, chatID int64, message string) {
	isAuth, user := db.IsAuth(chatID)

	if !isAuth {
		LanguageMsg(bot, chatID)
		return
	}

	if user.Location == "" {
		msg := tgbotapi.NewMessage(chatID, l.Language[user.Lang]["city"])
		msg.ReplyMarkup = helpKeyboard()
		_, err := bot.Send(msg)
		errors.CheckErrPanic(err)

		return
	}

	var (
		wthr string
		err  error
	)

	switch {
	case (message == "now") || (message == "/now") || (message == "сейчас"):
		wthr, err = w.CurrentWeather(user.Lat, user.Lng, user.Lang, user.Location)
		errors.CheckErrPanic(err)

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

	msg := tgbotapi.NewMessage(chatID, wthr)
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err = bot.Send(msg)
	errors.CheckErrPanic(err)
}
