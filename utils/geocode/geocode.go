package geocode

import (
	e "errors"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/google-geocode/geocode"
	c "github.com/lavrs/telegram-weather-bot/config"
	l "github.com/lavrs/telegram-weather-bot/language"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	"strconv"
)

func GetReverseGeocode(location *tgbotapi.Location, lang string) (*geocode.Geocoding, error) {
	var (
		g   *geocode.Geocoding
		err error
	)

	if lang == "ru" {
		g, err = geocode.ReverseGeocode(strconv.FormatFloat(
			location.Latitude, 'f', 6, 64),

			strconv.FormatFloat(location.Longitude, 'f', 6, 64),
			geocode.Russian, c.Cfg.GoogleGeocodeToken)
	} else {
		g, err = geocode.ReverseGeocode(strconv.FormatFloat(
			location.Latitude, 'f', 6, 64),

			strconv.FormatFloat(location.Longitude, 'f', 6, 64),
			geocode.English, c.Cfg.GoogleGeocodeToken)
	}
	if err != nil {
		if err.Error() == "ZERO_RESULTS" {
			return nil, e.New("_" + l.Language[lang]["ZERO_RESULTS_LOCATION"] + "_")
		}
		errors.CheckErrPanic(err)
	}

	return g, nil
}

func GetGeocode(location, lang string) (*geocode.Geocoding, error) {
	var (
		g   *geocode.Geocoding
		err error
	)

	if lang == "ru" {
		g, err = geocode.Geocode(location, geocode.Russian, c.Cfg.GoogleGeocodeToken)
	} else {
		g, err = geocode.Geocode(location, geocode.English, c.Cfg.GoogleGeocodeToken)
	}
	if err != nil {
		if err.Error() == "ZERO_RESULTS" {
			return nil, e.New("_" + l.Language[lang]["ZERO_RESULTS_CITY"] + "_")
		} else if err.Error() == "INVALID_REQUEST" {
			return nil, e.New("_" + l.Language[lang]["INVALID_REQUEST"] + "_")
		}

		errors.CheckErrPanic(err)
	}

	return g, nil
}
