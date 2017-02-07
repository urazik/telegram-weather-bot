package db

import (
	m "github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	r "gopkg.in/gorethink/gorethink.v3"
)

func createTelegramDB() {
	_, err := r.DBCreate("telegram").RunWrite(session)
	errors.CheckErrPanic(err)
}

func createUsersTable() {
	_, err := r.TableCreate("users").RunWrite(session)
	errors.CheckErrPanic(err)
}

func decodeOneBoolQueryResult(c *r.Cursor) (bool, error) {
	var res bool
	err := c.One(&res)
	if err != nil {
		return false, err
	}

	return res, nil
}

func getUser(telegramID int64) *m.DB {
	res, err := r.Table("users").Filter(
		r.Row.Field("telegramID").Eq(telegramID)).Run(session)
	errors.CheckErrPanic(err)
	defer res.Close()

	if res.IsNil() {
		return nil
	}

	var user m.DB
	err = res.One(&user)
	errors.CheckErrPanic(err)

	return &user
}

func getUserID(telegramID int64) *string {
	res, err := r.Table("users").Filter(
		r.Row.Field("telegramID").Eq(telegramID)).Field("id").Run(session)
	errors.CheckErrPanic(err)
	defer res.Close()

	if res.IsNil() {
		return nil
	}

	var ID string
	err = res.One(&ID)
	errors.CheckErrPanic(err)

	return &ID
}

func isTableAndDB() {
	query, err := r.DBList().Contains("telegram").Run(session)
	errors.CheckErrPanic(err)

	isDB, err := decodeOneBoolQueryResult(query)
	errors.CheckErrPanic(err)

	if isDB {
		query, err = r.TableList().Contains("users").Run(session)
		errors.CheckErrPanic(err)

		table, err := decodeOneBoolQueryResult(query)
		errors.CheckErrPanic(err)

		if !table {
			createUsersTable()
			return
		}

		return
	}

	createTelegramDB()
	createUsersTable()
	return
}
