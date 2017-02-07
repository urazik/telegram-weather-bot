package language

import "github.com/lavrs/telegram-weather-bot/model"

var (
	Language = map[string]map[string]string{
		"en": {
			"mps": "mps",
			"IFL": "feels like",
			"ZERO_RESULTS_LOCATION": "It is impossible to give a " +
				"forecast for the specified coordinates",
			"now":               "Now",
			"forToday":          "For today",
			"forTomorrow":       "For tomorrow",
			"forWeek":           "For week",
			"chooseLanguage":    "Choose language",
			"changeLanguageTo":  "Change language to",
			"ZERO_RESULTS_CITY": "Unable to find the weather for this city",
			"changeCityTo":      "Сhange city to",
			"emptycity":         "Enter your city to get the actual weather",
			"INVALID_REQUEST":   "The city is entered incorrectly",
			"Monday":            "Mon",
			"Tuesday":           "Tue",
			"Wednesday":         "Wed",
			"Thursday":          "Thu",
			"Friday":            "Fri",
			"Saturday":          "Sat",
			"Sunday":            "Sun",
			"YourLL":            "Your location and language",
			"empty_location":    "Empty location",

			"help": "_Bot_ show the *current* weather as well as *for today*, *tomorrow* and *the week*.\n\n" +
				"`Keyboard Description`:\n" +
				"`Now` - current weather\n" +
				"`For today` - weather for today\n" +
				"`For tomorrow` - weather for tomorrow\n" +
				"`For week` - weather for the week\n\n" +
				"You can also *send a message* to _bot_ or use a _bot_ *commands*.\n\n" +
				"You can select the language by *command* /lang or by " + model.GlobeWithMeridian + ".\n" +
				"View the current location and language is possible through " + model.Info + " or *command* /info.",
		},

		"ru": {
			"mps": "м/с",
			"IFL": "Чувствуется как",
			"ZERO_RESULTS_LOCATION": "Невозможно выдать погоду по данным координатам",
			"now":               "Сейчас",
			"forToday":          "На сегодня",
			"forTomorrow":       "На завтра",
			"forWeek":           "На неделю",
			"chooseLanguage":    "Выберите язык",
			"changeLanguageTo":  "Язык изменен на",
			"ZERO_RESULTS_CITY": "Невозможо выдать погоду в данном городе",
			"changeCityTo":      "Город изменен на",
			"emptycity":         "Введите город, чтобы получить актуальную погоду",
			"INVALID_REQUEST":   "Город введен неверно",
			"Monday":            "Пн",
			"Tuesday":           "Вт",
			"Wednesday":         "Ср",
			"Thursday":          "Чт",
			"Friday":            "Пт",
			"Saturday":          "Сб",
			"Sunday":            "Вс",
			"YourLL":            "Ваша локация и язык",
			"empty_location":    "Пустое местоположение",

			"help": "_Бот_ показывает *текущую* погоду, а также на *сегодня*, *завтра* и *неделю*.\n\n" +
				"`Описание клавиатуры`:\n" +
				"`Сейчас` - текущая погода\n" +
				"`На сегодня` - погода на сегодня\n" +
				"`На завтра` - погода на завтра\n" +
				"`На неделю` - погода на неделю\n\n" +
				"Также вы можете *отправить* _боту_ сообщение или использовать *команды*.\n\n" +
				"Вы можете выбрать язык введя *команду* /lang или " + model.GlobeWithMeridian + ".\n" +
				"Посмотреть текущие настройки локации и языка можно через " + model.Info + " или *команду* /info.",
		},
	}

	ChooseLanguage = "Choose language\nВыберите язык"
)
