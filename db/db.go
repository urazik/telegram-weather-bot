package db

import (
	"github.com/lavrs/google-geocode/geocode"
	c "github.com/lavrs/telegram-weather-bot/config"
	m "github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
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
	var (
		err error
		g   *geocode.Geocoding
	)

	ID := getUserID(telegramID)

	if user.Lang == lang {
		return lang
	} else if user.Location != "" {
		if lang == "ru" {
			g, err = geocode.Geocode(user.Location, geocode.Russian, c.Cfg.GoogleGeocodeToken)
		} else {
			g, err = geocode.Geocode(user.Location, geocode.English, c.Cfg.GoogleGeocodeToken)
		}
		errors.CheckErrPanic(err)

		var data = map[string]interface{}{
			"lang": lang,
			"city": g.Result[0].FormattedAddress,
		}

		_, err = r.Table("users").Get(ID).Update(data).RunWrite(session)
		errors.CheckErrPanic(err)

		return lang
	} else {
		var data = map[string]interface{}{
			"lang": lang,
		}

		_, err = r.Table("users").Get(ID).Update(data).RunWrite(session)
		errors.CheckErrPanic(err)

		return lang
	}
}

func updateUser(ID string, g *geocode.Geocoding) {
	var data = map[string]interface{}{
		"city": g.Result[0].FormattedAddress,
		"lat":  g.Result[0].Geometry.Location.Lat,
		"lng":  g.Result[0].Geometry.Location.Lng,
	}

	_, err := r.Table("users").Get(ID).Update(data).RunWrite(session)
	errors.CheckErrPanic(err)
}

func SetUser(telegramID int64, g *geocode.Geocoding, lang string) {
	userID := getUserID(telegramID)

	if userID != nil {
		updateUser(*userID, g)
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
			"city":       g.Result[0].FormattedAddress,
			"lang":       lang,
			"lat":        g.Result[0].Geometry.Location.Lat,
			"lng":        g.Result[0].Geometry.Location.Lng,
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
