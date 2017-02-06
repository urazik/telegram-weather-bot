package forecast

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable"`
	DarkSkyStations    []string `json:"darksky-stations"`
	DataPointStations  []string `json:"datapoint-stations"`
	ISDStations        []string `json:"isds-stations"`
	LAMPStations       []string `json:"lamp-stations"`
	METARStations      []string `json:"metars-stations"`
	METNOLicense       string   `json:"metnol-license"`
	Sources            []string `json:"sources"`
	Units              string   `json:"units"`
}

type DataPoint struct {
	Time                       int64   `json:"time"`
	Summary                    string  `json:"summary"`
	Icon                       string  `json:"icon"`
	SunriseTime                int64   `json:"sunriseTime"`
	SunsetTime                 int64   `json:"sunsetTime"`
	PrecipIntensity            float64 `json:"precipIntensity"`
	PrecipIntensityMax         float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime     int64   `json:"precipIntensityMaxTime"`
	PrecipProbability          float64 `json:"precipProbability"`
	PrecipType                 string  `json:"precipType"`
	PrecipAccumulation         float64 `json:"precipAccumulation"`
	Temperature                float64 `json:"temperature"`
	TemperatureMin             float64 `json:"temperatureMin"`
	TemperatureMinTime         int64   `json:"temperatureMinTime"`
	TemperatureMax             float64 `json:"temperatureMax"`
	TemperatureMaxTime         int64   `json:"temperatureMaxTime"`
	ApparentTemperature        float64 `json:"apparentTemperature"`
	ApparentTemperatureMax     float64 `json:"apparentTemperature"`
	ApparentTemperatureMaxTime int64   `json:"apparentTemperature"`
	ApparentTemperatureMin     float64 `json:"apparentTemperature"`
	ApparentTemperatureMinTime int64   `json:"apparentTemperature"`
	DewPoint                   float64 `json:"dewPoint"`
	WindSpeed                  float64 `json:"windSpeed"`
	WindBearing                float64 `json:"windBearing"`
	CloudCover                 float64 `json:"cloudCover"`
	Humidity                   float64 `json:"humidity"`
	Pressure                   float64 `json:"pressure"`
	Visibility                 float64 `json:"visibility"`
	Ozone                      float64 `json:"ozone"`
	MoonPhase                  float64 `json:"moonPhase"`
	NearestStormBearing        float64 `json:"nearestStormBearing"`
	NearestStormDistance       float64 `json:"nearestStormDistance"`
}

type DataBlock struct {
	Summary string      `json:"summary"`
	Icon    string      `json:"icon"`
	Data    []DataPoint `json:"data"`
}

type Alert struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Time        int64  `json:"time"`
	Expires     int64  `json:"expires"`
	URI         string `json:"uri"`
}

type Forecast struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timezone  string    `json:"timezone"`
	Currently DataPoint `json:"currently"`
	Minutely  DataBlock `json:"minutely"`
	Hourly    DataBlock `json:"hourly"`
	Daily     DataBlock `json:"daily"`
	Alerts    []Alert   `json:"alerts"`
	Flags     Flags     `json:"flags"`
}

type units string
type language string

const (
	BASEURL = "https://api.darksky.net/forecast"

	CA   units = "ca"
	SI   units = "si"
	US   units = "us"
	UK   units = "uk"
	AUTO units = "auto"

	Arabic             language = "ar"
	Azerbaijani        language = "az"
	Belarusian         language = "be"
	Bosnian            language = "bs"
	Catalan            language = "ca"
	Czech              language = "cs"
	German             language = "de"
	Greek              language = "el"
	English            language = "en"
	Spanish            language = "es"
	Estonian           language = "et"
	French             language = "fr"
	Croatian           language = "hr"
	Hungarian          language = "hu"
	Indonesian         language = "id"
	Italian            language = "it"
	Icelandic          language = "is"
	Cornish            language = "kw"
	NorwegianBokmal    language = "nb"
	Dutch              language = "nl"
	Polish             language = "pl"
	Portuguese         language = "pt"
	Russian            language = "ru"
	Slovak             language = "sk"
	Slovenian          language = "sl"
	Serbian            language = "sr"
	Swedish            language = "sv"
	Tetum              language = "tet"
	Turkish            language = "tr"
	Ukrainian          language = "uk"
	IgpayAtinlay       language = "x-pig-latin"
	Chinese            language = "zh"
	TraditionalChinese language = "zh-tw"
)

func GetForecast(token, lat, lng, time string, lang language, units units) (*Forecast, error) {
	res, err := forecastResponse(token, lat, lng, time, lang, units)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	f, err := fromJson(res)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func fromJson(res *http.Response) (*Forecast, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var f Forecast
	err = json.Unmarshal(body, &f)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func forecastResponse(token, lat, lng, time string, lang language, units units) (*http.Response, error) {
	coord := lat + "," + lng

	var url string
	if time == "now" {
		url = BASEURL + "/" + token + "/" + coord + "?lang=" + string(lang) + "&units=" + string(units)
	} else {
		url = BASEURL + "/" + token + "/" + coord + "," + time + "?lang=" + string(lang) + "&units=" + string(units)
	}

	res, err := http.Get(url)
	if err != nil {
		return res, err
	}

	return res, nil
}
