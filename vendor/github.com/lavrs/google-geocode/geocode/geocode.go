package geocode

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	u "net/url"
)

type Geocoding struct {
	Result []Result `json:"results"`
	Status string   `json:"status"`
}

type Result struct {
	AddressComponents []AddressComponents `json:"address_components"`
	FormattedAddress  string              `json:"formatted_address"`
	Geometry          Geometry            `json:"geometry"`
	PlaceID           string              `json:"place_id"`
	Types             []string            `json:"types"`
}

type Geometry struct {
	Location     Location       `json:"location"`
	LocationType string         `json:"location_type"`
	Bounds       ViewportBounds `json:"bounds"`
	Viewport     ViewportBounds `json:"viewport"`
}

type ViewportBounds struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type AddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type language string

const (
	BASEURL = "https://maps.googleapis.com/maps/api/geocode/json"

	Arabic              language = "ar"
	Azerbaijani         language = "az"
	Belarusian          language = "be"
	Bosnian             language = "bs"
	Catalan             language = "ca"
	Czech               language = "cs"
	German              language = "de"
	Greek               language = "el"
	English             language = "en"
	Spanish             language = "es"
	Estonian            language = "et"
	French              language = "fr"
	Croatian            language = "hr"
	Hungarian           language = "hu"
	Indonesian          language = "id"
	Italian             language = "it"
	Icelandic           language = "is"
	Cornish             language = "kw"
	NorwegianBokmal     language = "nb"
	Dutch               language = "nl"
	Polish              language = "pl"
	Portuguese          language = "pt"
	Russian             language = "ru"
	Slovak              language = "sk"
	Slovenian           language = "sl"
	Serbian             language = "sr"
	Swedish             language = "sv"
	Tetum               language = "tet"
	Turkish             language = "tr"
	Ukrainian           language = "uk"
	IgpayAtinlay        language = "x-pig-latin"
	Chinese             language = "zh"
	TraditionalChinese  language = "zh-tw"
	Kannada             language = "kn"
	Bulgarian           language = "bg"
	Korean              language = "ko"
	Bengali             language = "bn"
	Lithuanian          language = "lt"
	EnglishAustralian   language = "en-AU"
	Norwegian           language = "no"
	Latvian             language = "lv"
	Marathi             language = "mr"
	Danish              language = "da"
	Malayalam           language = "ml"
	EnglishGreatBritain language = "en-GB"
	PortugueseBrazil    language = "pt-BR"
	PortuguesePortugal  language = "pt-PT"
	Basque              language = "eu"
	Romanian            language = "ro"
	Farsi               language = "fa"
	Finnish             language = "fi"
	Hebrew              language = "iw"
	Vietnamese          language = "vi"
	Tagalog             language = "tl"
	Thai                language = "th"
	Hindi               language = "hi"
	Telugu              language = "ta"
	Filipino            language = "fil"
	Gujarati            language = "gu"
	Tamil               language = "ta"
	Galician            language = "gl"
	Japanese            language = "ja"
)

func Geocode(city string, lang language, token string) (*Geocoding, error) {
	res, err := geocodeResponse(city, lang, token)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	g, err := fromJson(res)
	if err != nil {
		return nil, err
	}

	if g.Status != "OK" {
		return nil, errors.New(g.Status)
	}

	return g, nil
}

func geocodeResponse(city string, lang language, token string) (*http.Response, error) {
	url := BASEURL + "?address=" + u.QueryEscape(city) + "&language=" + string(lang) + "&key=" + token
	res, err := http.Get(url)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ReverseGeocode(lat, lng string, lang language, token string) (*Geocoding, error) {
	res, err := reverseGeocodeResponse(lat, lng, lang, token)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	g, err := fromJson(res)
	if err != nil {
		return nil, err
	}

	if g.Status != "OK" {
		return nil, errors.New(g.Status)
	}

	return g, nil
}

func fromJson(res *http.Response) (*Geocoding, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var g Geocoding
	err = json.Unmarshal(body, &g)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func reverseGeocodeResponse(lat, lng string, lang language, token string) (*http.Response, error) {
	latlng := lat + "," + lng

	url := BASEURL + "?latlng=" + latlng + "&language=" + string(lang) + "&key=" + token
	res, err := http.Get(url)
	if err != nil {
		return res, err
	}

	return res, nil
}
