package weather

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/darksky/forecast"
	"github.com/lavrs/google-geocode/geocode"
	c "github.com/lavrs/telegram-weather-bot/config"
	l "github.com/lavrs/telegram-weather-bot/language"
	"github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	"math"
	"strconv"
	"time"
)

func FTS0(i float64) string {
	return strconv.FormatFloat(i, 'f', 0, 64)
}

func FTS6(i float64) string {
	return strconv.FormatFloat(i, 'f', 6, 64)
}

func getReverseGeocoding(location *tgbotapi.Location, lang string) (*geocode.Geocoding, error) {
	var (
		g   *geocode.Geocoding
		err error
	)

	if lang == "ru" {
		g, err = geocode.ReverseGeocode(strconv.FormatFloat(location.Latitude, 'f', 6, 64),
			strconv.FormatFloat(location.Longitude, 'f', 6, 64),
			geocode.Russian, c.Cfg.GoogleGeocodeToken)
	} else {
		g, err = geocode.ReverseGeocode(strconv.FormatFloat(location.Latitude, 'f', 6, 64),
			strconv.FormatFloat(location.Longitude, 'f', 6, 64),
			geocode.English, c.Cfg.GoogleGeocodeToken)
	}
	if err != nil {
		return nil, err
	}

	return g, nil
}

func getWeatherByDay(user *model.DB, f forecast.DataPoint, timezone string) string {
	return getDate(f.Time, timezone, user.Lang) + "," + getCity(user.Location) + "\n`" + f.Summary + "`\n\n" +
		model.Icons[f.Icon] + " *" + FTS0(f.TemperatureMin) + ".." + FTS0(f.TemperatureMax) + "°С*  " +
		model.Directions[int(math.Mod(f.WindBearing/22.5+.5, 16))] + " *" + FTS0(f.WindSpeed) + " " + l.Language[user.Lang]["mps"] + "*\n" +
		model.Sunrise + " " + getTime(f.SunriseTime, timezone) + "  " +
		model.Sunset + " " + getTime(f.SunsetTime, timezone) + "  " + model.Moons[getMoonPhase(f.MoonPhase)] + "\n" +
		"`" + l.Language[user.Lang]["IFL"] + "`  *" + FTS0(f.ApparentTemperatureMin) + ".." + FTS0(f.ApparentTemperatureMax) + "°C*"
}

func getWeekWeather(user *model.DB, f *forecast.Forecast) string {
	var text string

	text = "`" + user.Location + "`\n\n`" + f.Daily.Summary + "`\n\n"
	for _, day := range f.Daily.Data {
		text += getDate(day.Time, f.Timezone, user.Lang) + "  " +
			model.Icons[day.Icon] + " *" + FTS0(day.TemperatureMin) + ".." + FTS0(day.TemperatureMax) + "°С*  " +
			model.Directions[int(math.Mod(day.WindBearing/22.5+.5, 16))] + " *" + FTS0(day.WindSpeed) + " " + l.Language[user.Lang]["mps"] + "*\n" +
			"`" + day.Summary + "`\n\n"
	}

	return text
}

func getMoonPhase(phase float64) string {
	if phase < 0.25 {
		return "new moon"
	} else if phase < 0.50 {
		return "first quarter moon"
	} else if phase < 0.75 {
		return "full moon"
	} else {
		return "last quarter moon"
	}
}

func getForecast(lat, lng float64, lang string) *forecast.Forecast {
	var (
		f   *forecast.Forecast
		err error
	)

	if lang == "ru" {
		f, err = forecast.GetForecast(c.Cfg.DarkskyToken, FTS6(lat), FTS6(lng), "now", forecast.Russian, forecast.SI)
	} else {
		f, err = forecast.GetForecast(c.Cfg.DarkskyToken, FTS6(lat), FTS6(lng), "now", forecast.English, forecast.SI)
	}
	errors.CheckErrPanic(err)

	return f
}

func getCurrentWeather(lang string, f *forecast.Forecast) string {
	return model.Icons[f.Currently.Icon] + " *" + FTS0(f.Currently.Temperature) + "°С*  " +
		model.Directions[int(math.Mod(f.Currently.WindBearing/22.5+.5, 16))] + " *" +
		FTS0(f.Currently.WindSpeed) + " " + l.Language[lang]["mps"] + "*  `" + f.Currently.Summary + ".`\n" +
		"`" + l.Language[lang]["IFL"] + "`  *" + FTS0(f.Currently.ApparentTemperature) + "°C*"
}

func getCity(city string) string {
	return "   `" + city + "`\n"
}

func getTime(ftime int64, timezone string) string {
	time := getLocalTime(ftime, timezone)

	return "_" + time[11:16] + "_"
}

func getDate(ftime int64, timezone, lang string) string {
	date := getLocalTime(ftime, timezone)

	return "_" + date[8:10] + "/" + date[5:7] + " " + getWeekday(ftime, timezone, lang) + "_"
}

func getLocalTime(ftime int64, ftimezone string) string {
	timezone, err := time.LoadLocation(ftimezone)
	errors.CheckErrPanic(err)

	return time.Unix(int64(ftime), 0).In(timezone).String()
}

func getWeekday(ftime int64, ftimezone, lang string) string {
	timezone, err := time.LoadLocation(ftimezone)
	errors.CheckErrPanic(err)

	return l.Language[lang][time.Unix(int64(ftime), 0).In(timezone).Weekday().String()]
}
