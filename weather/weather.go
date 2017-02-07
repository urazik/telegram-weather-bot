package weather

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/telegram-weather-bot/model"
)

func CurrentWeather(lat, lng float64, lang, location string) string {
	f := getForecast(lat, lng, lang)

	return getTime(f.Currently.Time, f.Timezone) + getCity(location) + getCurrentWeather(lang, f)
}

func CurrentWeatherFromLocation(lang string, coord *tgbotapi.Location, location string) string {
	f := getForecast(coord.Latitude, coord.Longitude, lang)

	return getTime(f.Currently.Time, f.Timezone) + getCity(location) + getCurrentWeather(lang, f)
}

func WeatherOfDay(user *model.DB) (string, error) {
	f := getForecast(user.Lat, user.Lng, user.Lang)

	return getWeatherByDay(user, f.Daily.Data[0], f.Timezone), nil
}

func TomorrowWeather(user *model.DB) (string, error) {
	f := getForecast(user.Lat, user.Lng, user.Lang)

	return getWeatherByDay(user, f.Daily.Data[1], f.Timezone), nil
}

func WeekWeather(user *model.DB) (string, error) {
	f := getForecast(user.Lat, user.Lng, user.Lang)

	return getWeekWeather(user, f), nil
}
