package db

import (
    m "github.com/lavrs/telegram-weather-bot/model"
    "github.com/lavrs/telegram-weather-bot/utils/errors"
    r "gopkg.in/gorethink/gorethink.v3"
)

func createTelegramDB() {
    _, err := r.DBCreate("telegram").RunWrite(session)
    errors.Check(err)
}

func createUsersTable() {
    _, err := r.TableCreate("users").RunWrite(session)
    errors.Check(err)
}

func decodeOneBoolQueryResult(c *r.Cursor) (bool, error) {
    var res bool
    if err := c.One(&res); err != nil {
        return false, err
    }

    return res, nil
}

func getUser(telegramID int64) *m.DB {
    res, err := r.Table("users").Filter(
        r.Row.Field("telegramID").Eq(telegramID)).Run(session)
    errors.Check(err)
    defer res.Close()

    if res.IsNil() {
        return nil
    }

    var user m.DB
    err = res.One(&user)
    errors.Check(err)

    return &user
}

func getUserID(telegramID int64) *string {
    res, err := r.Table("users").Filter(
        r.Row.Field("telegramID").Eq(telegramID)).Field("id").Run(session)
    errors.Check(err)
    defer res.Close()

    if res.IsNil() {
        return nil
    }

    var ID string
    err = res.One(&ID)
    errors.Check(err)

    return &ID
}

func isTableAndDB() {
    query, err := r.DBList().Contains("telegram").Run(session)
    errors.Check(err)

    isDB, err := decodeOneBoolQueryResult(query)
    errors.Check(err)

    if isDB {
        query, err = r.TableList().Contains("users").Run(session)
        errors.Check(err)

        table, err := decodeOneBoolQueryResult(query)
        errors.Check(err)

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
