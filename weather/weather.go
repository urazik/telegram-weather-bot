package weather

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	l "github.com/lavrs/telegram-weather-bot/language"
	"github.com/lavrs/telegram-weather-bot/model"
)

func CurrentWeather(lat, lng float64, lang, city string) (string, error) {
	f := getForecast(lat, lng, lang)

	return getTime(f.Currently.Time, f.Timezone) + getCity(city) + getCurrentWeather(lang, f), nil
}

func CurrentWeatherFromLocation(lang string, location *tgbotapi.Location) (string, error) {
	g, err := getReverseGeocoding(location, lang)
	if err != nil {
		if err.Error() == "ZERO_RESULTS" {
			return "_" + l.Language[lang]["ZERO_RESULTS_LOCATION"] + "_", nil
		}

		return "", err
	}

	f := getForecast(location.Latitude, location.Longitude, lang)

	return getTime(f.Currently.Time, f.Timezone) + getCity(g.Result[0].FormattedAddress) + getCurrentWeather(lang, f), nil
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
