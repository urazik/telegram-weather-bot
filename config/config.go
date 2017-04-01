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

    &Cfg = data
}

func openFile() *m.Config {
    var (
        file []byte
        err  error
    )

    if file, err = ioutil.ReadFile("config.yml"); err != nil {
        log.Panic(err)
    }

    var data m.Config
    if err := yaml.Unmarshal(file, &data); err != nil {
        log.Panic(err)
    }

    return &data
}
