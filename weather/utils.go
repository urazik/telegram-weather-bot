package weather

import (
	"github.com/lavrs/darksky/forecast"
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

func getWeatherByDay(user *model.DB, f forecast.DataPoint, timezone string) string {
	return getDate(f.Time, timezone, user.Lang) + "," + getCity(user.Location) +
		"\n`" + f.Summary + "`\n\n" + model.Icons[f.Icon] + " *" +
		FTS0(f.TemperatureMin) + ".." + FTS0(f.TemperatureMax) + getTempUnit(user.Units) + "*" +
		"  *" + getWind(f.WindSpeed, f.WindBearing, user.Lang, user.Units) +
		"* \n" + model.Sunrise + " " + getTime(f.SunriseTime, timezone) +
		"  " + model.Sunset + " " + getTime(f.SunsetTime, timezone) +
		"  " + model.Moons[getMoonPhase(f.MoonPhase)] + "\n" +
		"`" + l.Language[user.Lang]["IFL"] + "`  *" +
		FTS0(f.ApparentTemperatureMin) + ".." + FTS0(f.ApparentTemperatureMax) + getTempUnit(user.Units) + "*"
}

func getWeekWeather(user *model.DB, f *forecast.Forecast) string {
	var text string

	text = "`" + user.Location + "`\n\n`" + f.Daily.Summary + "`\n\n"
	for _, day := range f.Daily.Data {
		text += getDate(day.Time, f.Timezone, user.Lang) + "  " +
			model.Icons[day.Icon] + " *" + FTS0(day.TemperatureMin) +
			".." + FTS0(day.TemperatureMax) + getTempUnit(user.Units) + "*" +
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

func getForecast(lat, lng float64, lang, units string) *forecast.Forecast {
	var (
		f   *forecast.Forecast
		err error
	)

	if lang == "ru" {
		if units == "si" {
			f, err = forecast.GetForecast(
				c.Cfg.DarkskyToken, FTS6(lat), FTS6(lng), "now", forecast.Russian, forecast.SI)
		} else {
			f, err = forecast.GetForecast(
				c.Cfg.DarkskyToken, FTS6(lat), FTS6(lng), "now", forecast.Russian, forecast.US)
		}
	} else {
		if units == "si" {
			f, err = forecast.GetForecast(
				c.Cfg.DarkskyToken, FTS6(lat), FTS6(lng), "now", forecast.English, forecast.SI)
		} else {
			f, err = forecast.GetForecast(
				c.Cfg.DarkskyToken, FTS6(lat), FTS6(lng), "now", forecast.English, forecast.US)
		}
	}
	errors.CheckErrPanic(err)

	if f.APICalls == 995 {
		return nil
	}

	return f
}

func getCurrentWeather(lang string, units string, f *forecast.Forecast) string {
	return model.Icons[f.Currently.Icon] + " *" + FTS0(f.Currently.Temperature) +
		getTempUnit(units) + "  " +
		getWind(f.Currently.WindSpeed, f.Currently.WindBearing, lang, units) +
		"*  `" + f.Currently.Summary + ".`\n`" + l.Language[lang]["IFL"] +
		"`  *" + FTS0(f.Currently.ApparentTemperature) + getTempUnit(units) + "*"
}

func getWind(speed, bearing float64, lang, units string) string {
	return model.Directions[int(math.Mod(360+bearing/22.5+.5, 16))] +
		" " + FTS0(speed) + " " + getWindUnit(lang, units)
}

func getTempUnit(units string) string {
	if units == "si" {
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
	time := getLocalTime(ftime, timezone)

	return "_" + time[11:16] + "_"
}

func getDate(ftime int64, timezone, lang string) string {
	date := getLocalTime(ftime, timezone)

	return "_" + date[8:10] + "/" + date[5:7] +
		" " + getWeekday(ftime, timezone, lang) + "_"
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
