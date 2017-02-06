package msg

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/telegram-weather-bot/model"
	"strings"
)

func Updates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	update.Message.Text = strings.ToLower(update.Message.Text)

	if update.Message.Command() == "start" {
		StartMsg(bot, update.Message.Chat.ID)

		return
	}

	if (update.Message.Text == "now") || (update.Message.Text == "for today") ||
		(update.Message.Text == "for tomorrow") || (update.Message.Text == "for week") ||
		(update.Message.Text == "сейчас") || (update.Message.Text == "на сегодня") ||
		(update.Message.Text == "на завтра") || (update.Message.Text == "на неделю") ||
		(update.Message.Command() == "now") || (update.Message.Command() == "today") ||
		(update.Message.Command() == "tomorrow") || (update.Message.Command() == "week") {

		WeatherMsgFromCmd(bot, update.Message.Chat.ID, update.Message.Text)

		return
	}

	if (update.Message.Text == model.CountriesFATE["en"]) ||
		(update.Message.Text == model.CountriesFATE["ru"]) {

		UpdateLangMsg(bot, update.Message.Chat.ID, update.Message.Text)

		return
	}

	if (update.Message.Text == model.Info) || (update.Message.Command() == "info") {
		InfoMsg(bot, update.Message.Chat.ID)

		return
	}

	if (update.Message.Text == model.GlobeWithMeridian) || (update.Message.Command() == "lang") {
		LanguageMsg(bot, update.Message.Chat.ID)

		return
	}

	if (update.Message.Command() == "help") || (update.Message.Text == model.Help) {
		Help(bot, update.Message.Chat.ID)

		return
	}

	if update.Message.Location != nil {
		WeatherMsgFromLocation(bot, update.Message.Chat.ID, update.Message.Location)

		return
	}

	WeatherMsgFromCity(bot, update.Message.Chat.ID, update.Message.Text)
}
