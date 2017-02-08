package weather

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/telegram-weather-bot/model"
)

func CurrentWeather(lat, lng float64, lang, location, units string) string {
	f := getForecast(lat, lng, lang, units)

	return getTime(f.Currently.Time, f.Timezone) + getCity(location) + getCurrentWeather(lang, units, f)
}

func CurrentWeatherFromLocation(lang string, coord *tgbotapi.Location, location, units string) string {
	f := getForecast(coord.Latitude, coord.Longitude, lang, units)

	return getTime(f.Currently.Time, f.Timezone) + getCity(location) + getCurrentWeather(lang, units, f)
}

func WeatherOfDay(user *model.DB) string {
	f := getForecast(user.Lat, user.Lng, user.Lang, user.Units)

	return getWeatherByDay(user, f.Daily.Data[0], f.Timezone, user.Units)
}

func TomorrowWeather(user *model.DB) string {
	f := getForecast(user.Lat, user.Lng, user.Lang, user.Units)

	return getWeatherByDay(user, f.Daily.Data[1], f.Timezone, user.Units)
}

func WeekWeather(user *model.DB) string {
	f := getForecast(user.Lat, user.Lng, user.Lang, user.Units)

	return getWeekWeather(user, f)
}
