package db

import (
	gmodel "github.com/lavrs/google-geocode/geocode"
	m "github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	"github.com/lavrs/telegram-weather-bot/utils/geocode"
	r "gopkg.in/gorethink/gorethink.v3"
	"log"
)

var session *r.Session

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: "telegram",
	})
	if err != nil {
		log.Panic(err)
	}

	isTableAndDB()
}

func UpdateUserLang(user *m.DB, lang string, telegramID int64) string {
	ID := getUserID(telegramID)

	if user.Lang == lang {
		return lang
	} else if user.Location != "" {
		g, _ := geocode.GetGeocode(user.Location, lang)

		var data = map[string]interface{}{
			"lang":     lang,
			"location": g.Result[0].FormattedAddress,
		}

		_, err := r.Table("users").Get(ID).Update(data).RunWrite(session)
		errors.CheckErrPanic(err)

		return lang
	} else {
		var data = map[string]interface{}{
			"lang": lang,
		}

		_, err := r.Table("users").Get(ID).Update(data).RunWrite(session)
		errors.CheckErrPanic(err)

		return lang
	}
}

func UpdateUserUnits(telegramID int64, units string) {
	ID := getUserID(telegramID)

	if (units == "°c, mps") || (units == "°c, м/c") {
		units = "si"
	} else {
		units = "us"
	}

	var data = map[string]interface{}{
		"units": units,
	}

	_, err := r.Table("users").Get(ID).Update(data).RunWrite(session)
	errors.CheckErrPanic(err)
}

func updateUserLocation(ID string, g *gmodel.Geocoding) {
	var data = map[string]interface{}{
		"location": g.Result[0].FormattedAddress,
		"lat":      g.Result[0].Geometry.Location.Lat,
		"lng":      g.Result[0].Geometry.Location.Lng,
	}

	_, err := r.Table("users").Get(ID).Update(data).RunWrite(session)
	errors.CheckErrPanic(err)
}

func SetUser(telegramID int64, g *gmodel.Geocoding, lang string) {
	userID := getUserID(telegramID)

	if userID != nil {
		updateUserLocation(*userID, g)
		return
	}

	var data = map[string]interface{}{}

	if g == nil {
		data = map[string]interface{}{
			"telegramID": telegramID,
			"lang":       lang,
		}
	} else {
		data = map[string]interface{}{
			"telegramID": telegramID,
			"location":   g.Result[0].FormattedAddress,
			"lang":       lang,
			"lat":        g.Result[0].Geometry.Location.Lat,
			"lng":        g.Result[0].Geometry.Location.Lng,
			"units":      "si",
		}
	}

	_, err := r.Table("users").Insert(data).RunWrite(session)
	errors.CheckErrPanic(err)
}

func IsAuth(telegramID int64) (bool, *m.DB) {
	user := getUser(telegramID)

	if user == nil {
		return false, nil
	} else {
		return true, user
	}
}
