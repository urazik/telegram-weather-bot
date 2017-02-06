package config

import (
	"github.com/go-yaml/yaml"
	m "github.com/lavrs/telegram-weather-bot/model"
	"io/ioutil"
	"log"
)

var Cfg m.Config

func SetConfig() {
	data := openFile()

	Cfg.TelegramToken = data.TelegramToken
	Cfg.TelegramTestToken = data.TelegramTestToken
	Cfg.DarkskyToken = data.DarkskyToken
	Cfg.GoogleGeocodeToken = data.GoogleGeocodeToken
	Cfg.MyTelegramID = data.MyTelegramID
}

func openFile() *m.Config {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Panic(err)
	}

	var data m.Config
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Panic(err)
	}

	return &data
}
