package aqicn

import (
	"testing"
)

func TestGetWeather(t *testing.T) {
	info, e := GetWeatherInfo("Your Token", "shanghai")
	if e == nil {
		t.Log("OK")
		t.Log(info)

	} else {
		t.Fatal(e)
	}
}
