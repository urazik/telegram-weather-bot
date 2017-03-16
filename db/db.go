package db

import (
	"github.com/mlbright/forecast/v2"
	m "github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	"github.com/lavrs/telegram-weather-bot/utils/geocoding"
	"googlemaps.github.io/maps"
	r "gopkg.in/gorethink/gorethink.v3"
	"log"
)

var session *r.Session

func init() {
	var err error

	if session, err = r.Connect(r.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: "telegram",
	}); err != nil {
		log.Panic(err)
	}

	isTableAndDB()
}

func UpdateUserLang(user *m.DB, lang string, telegramID int64) string {
	ID := getUserID(telegramID)

	if user.Lang == lang {
		return lang
	} else if user.Location != "" {
		g, err := geocoding.Geocode(user.Location, lang)
		errors.Check(err)

		var data = map[string]interface{}{
			"lang":     lang,
			"location": g[0].FormattedAddress,
		}

		_, err = r.Table("users").Get(ID).Update(data).RunWrite(session)
		errors.Check(err)

		return lang
	} else {
		var data = map[string]interface{}{
			"lang": lang,
		}

		_, err := r.Table("users").Get(ID).Update(data).RunWrite(session)
		errors.Check(err)

		return lang
	}
}

func UpdateUserUnits(telegramID int64, units string) {
	ID := getUserID(telegramID)

	if (units == "°c, mps") || (units == "°c, м/c") {
		units = string(forecast.SI)
	} else {
		units = string(forecast.US)
	}

	var data = map[string]interface{}{
		"units": units,
	}

	_, err := r.Table("users").Get(ID).Update(data).RunWrite(session)
	errors.Check(err)
}

func updateUserLocation(ID string, g []maps.GeocodingResult) {
	var data = map[string]interface{}{
		"location": g[0].FormattedAddress,
		"lat":      g[0].Geometry.Location.Lat,
		"lng":      g[0].Geometry.Location.Lng,
	}

	_, err := r.Table("users").Get(ID).Update(data).RunWrite(session)
	errors.Check(err)
}

func SetUser(telegramID int64, g []maps.GeocodingResult, lang string) {
	userID := getUserID(telegramID)

	if userID != nil {
		updateUserLocation(*userID, g)
		return
	}

	var data = map[string]interface{}{}
	data = map[string]interface{}{
		"telegramID": telegramID,
		"lang":       lang,
		"units":      forecast.SI,
	}

	_, err := r.Table("users").Insert(data).RunWrite(session)
	errors.Check(err)
}

func IsAuth(telegramID int64) (bool, *m.DB) {
	user := getUser(telegramID)

	if user == nil {
		return false, nil
	} else {
		return true, user
	}
}
