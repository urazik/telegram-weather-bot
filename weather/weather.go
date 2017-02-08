package weather

import (
	"github.com/lavrs/telegram-weather-bot/language"
	"github.com/lavrs/telegram-weather-bot/model"
)

func CurrentWeather(lat, lng float64, location string, user *model.DB) string {
	f := getForecast(lat, lng, user.Lang, user.Units)
	if f == nil {
		return "_" + language.Language[user.Lang]["unknownError"] + "_"
	}

	return getTime(f.Currently.Time, f.Timezone) +
		getCity(location) + getCurrentWeather(user.Lang, user.Units, f)
}

func CurrentWeatherFromLocation(lat, lng float64, location string, user *model.DB) string {
	f := getForecast(lat, lng, user.Lang, user.Units)
	if f == nil {
		return "_" + language.Language[user.Lang]["unknownError"] + "_"
	}

	return getTime(f.Currently.Time, f.Timezone) +
		getCity(location) + getCurrentWeather(user.Lang, user.Units, f)
}

func WeatherOfDay(user *model.DB) string {
	f := getForecast(user.Lat, user.Lng, user.Lang, user.Units)
	if f == nil {
		return "_" + language.Language[user.Lang]["unknownError"] + "_"
	}

	return getWeatherByDay(user, f.Daily.Data[0], f.Timezone)
}

func TomorrowWeather(user *model.DB) string {
	f := getForecast(user.Lat, user.Lng, user.Lang, user.Units)
	if f == nil {
		return "_" + language.Language[user.Lang]["unknownError"] + "_"
	}

	return getWeatherByDay(user, f.Daily.Data[1], f.Timezone)
}

func WeekWeather(user *model.DB) string {
	f := getForecast(user.Lat, user.Lng, user.Lang, user.Units)
	if f == nil {
		return "_" + language.Language[user.Lang]["unknownError"] + "_"
	}

	return getWeekWeather(user, f)
}
