package weather

import (
	l "github.com/lavrs/telegram-weather-bot/language"
	"github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	u "github.com/lavrs/telegram-weather-bot/utils/format"
	"github.com/mlbright/forecast/v2"
	"math"
	"time"
)

func getWeatherByDay(user *model.DB, f forecast.DataPoint, timezone string) string {
	return getDate(f.Time, timezone, user.Lang) + "," + getCity(user.Location) +
		"\n`" + f.Summary + "`\n\n" + model.Icons[f.Icon] + " *" +
		u.FTS0(f.TemperatureMin) + ".." + u.FTS0(f.TemperatureMax) + getTempUnit(user.Units) + "*" +
		"  *" + getWind(f.WindSpeed, f.WindBearing, user.Lang, user.Units) +
		"* \n" + model.Sunrise + " " + getTime(f.SunriseTime, timezone) +
		"  " + model.Sunset + " " + getTime(f.SunsetTime, timezone) +
		"  " + model.Moons[getMoonPhase(f.MoonPhase)] + "\n" +
		"`" + l.Language[user.Lang]["IFL"] + "`  *" +
		u.FTS0(f.ApparentTemperatureMin) + ".." + u.FTS0(f.ApparentTemperatureMax) + getTempUnit(user.Units) + "*"
}

func getWeekWeather(user *model.DB, f *forecast.Forecast) string {
	var text string

	text = "`" + user.Location + "`\n\n`" + f.Daily.Summary + "`\n\n"
	for _, day := range f.Daily.Data {
		text += getDate(day.Time, f.Timezone, user.Lang) + "  " +
			model.Icons[day.Icon] + " *" + u.FTS0(day.TemperatureMin) +
			".." + u.FTS0(day.TemperatureMax) + getTempUnit(user.Units) + "*" +
			"  *" + getWind(day.WindSpeed, day.WindBearing, user.Lang, user.Units) +
			"*\n`" + day.Summary + "`\n\n"
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

func getCurrentWeather(lang string, units string, f *forecast.Forecast) string {
	return model.Icons[f.Currently.Icon] + " *" + u.FTS0(f.Currently.Temperature) +
		getTempUnit(units) + "  " +
		getWind(f.Currently.WindSpeed, f.Currently.WindBearing, lang, units) +
		"*  `" + f.Currently.Summary + ".`\n`" + l.Language[lang]["IFL"] +
		"`  *" + u.FTS0(f.Currently.ApparentTemperature) + getTempUnit(units) + "*"
}

func getWind(speed, bearing float64, lang, units string) string {
	return model.Directions[int(math.Mod(360+bearing/22.5+.5, 16))] +
		" " + u.FTS0(speed) + " " + getWindUnit(lang, units)
}

func getTempUnit(units string) string {
	if units == string(forecast.SI) {
		return "°С"
	} else {
		return "°F"
	}
}

func getWindUnit(lang, units string) string {
	if units == "si" {
		return l.Language[lang]["mps"]
	} else {
		return l.Language[lang]["mph"]
	}
}

func getCity(city string) string {
	return "   `" + city + "`\n"
}

func getTime(ftime int64, timezone string) string {
	return "_" + getLocalTime(ftime, timezone)[11:16] + "_"
}

func getDate(ftime int64, timezone, lang string) string {
	date := getLocalTime(ftime, timezone)

	return "_" + date[8:10] + "/" + date[5:7] +
		" " + getWeekday(ftime, timezone, lang) + "_"
}

func getLocalTime(ftime int64, ftimezone string) string {
	timezone, err := time.LoadLocation(ftimezone)
	errors.Check(err)

	return time.Unix(int64(ftime), 0).In(timezone).String()
}

func getWeekday(ftime int64, ftimezone, lang string) string {
	timezone, err := time.LoadLocation(ftimezone)
	errors.Check(err)

	return l.Language[lang][time.Unix(int64(ftime), 0).In(timezone).Weekday().String()]
}
