package aqicn

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	AqiUrl string = "https://api.waqi.info/feed/"
)

type WeatherInfo struct {
	Status string `json:"status"`
	Data   struct {
		Aqi          int `json:"aqi"`
		Idx          int `json:"idx"`
		Attributions []struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"attributions"`
		City struct {
			Geo  []float64 `json:"geo"`
			Name string    `json:"name"`
			URL  string    `json:"url"`
		} `json:"city"`
		Dominentpol string `json:"dominentpol"`
		Iaqi        struct {
			Co struct {
				V float64 `json:"v"`
			} `json:"co"`
			H struct {
				V float64 `json:"v"`
			} `json:"h"`
			O3 struct {
				V float64 `json:"v"`
			} `json:"o3"`
			No2 struct {
				V float64 `json:"v"`
			} `json:"no2"`
			P struct {
				V float64 `json:"v"`
			} `json:"p"`
			Pm10 struct {
				V float64 `json:"v"`
			} `json:"pm10"`
			Pm25 struct {
				V float64 `json:"v"`
			} `json:"pm25"`
			So2 struct {
				V float64 `json:"v"`
			} `json:"so2"`
			T struct {
				V float64 `json:"v"`
			} `json:"t"`
			W struct {
				V float64 `json:"v"`
			} `json:"w"`
		} `json:"iaqi"`
		Time struct {
			S  string `json:"s"`
			Tz string `json:"tz"`
			V  int    `json:"v"`
		} `json:"time"`
		Debug struct {
			Sync time.Time `json:"sync"`
		} `json:"debug"`
	} `json:"data"`
}

func GetWeatherInfo(token string, city string) (*WeatherInfo, error) {
	resp, err := http.Get(AqiUrl + city + "/?token=" + token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseResp(bytes)
}

func parseResp(bytes []byte) (*WeatherInfo, error) {
	parsedJson := gjson.Parse(string(bytes))
	if parsedJson.Get("status").Value().(string) != "ok" {
		return nil, errors.New(parsedJson.Get("data").Value().(string))
	}

	weatherInfo := WeatherInfo{}
	err := json.Unmarshal(bytes, &weatherInfo)
	if err != nil {
		return nil, err
	}
	return &weatherInfo, nil
}
