package openweather

import (
	"app/logging"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Values struct {
	Name    string
	Main    Main
	Weather Weather
}

type response struct {
	Name    string
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
}

type Weather struct {
	Icon string `json:"icon"`
}

type OpenWeather struct {
	APIkey  string
	Celsius bool
	Debug   logging.LogLevel
}

func (weather OpenWeather) Get(city string) (Values, error) {
	logger := logging.Logger{
		Level: weather.Debug,
	}
	locationData, err := weather.fetchLocationData(city)

	if err != nil {
		logger.Debug(err.Error())
		return Values{}, errors.New("error fetching location data")
	}

	request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s",
		locationData.Lat, locationData.Lon, weather.APIkey)

	logger.Debug(request)

	resp, err := http.Get(request)

	if err != nil {
		logger.Debug(err.Error())
		return Values{}, errors.New("error fetching openweather data")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Debug(err.Error())
		return Values{}, errors.New("error reading response body")
	}

	var data response
	err = json.Unmarshal(body, &data)

	if err != nil {
		logger.Debug(err.Error())
		return Values{}, errors.New("unmarshal error")
	}

	if weather.Celsius {
		data.Main.Temp = kelvinToCelsius(data.Main.Temp)
		data.Main.TempMax = kelvinToCelsius(data.Main.TempMax)
		data.Main.TempMin = kelvinToCelsius(data.Main.TempMin)
		data.Main.FeelsLike = kelvinToCelsius(data.Main.FeelsLike)
	}

	values := Values{
		Name:    locationData.Name,
		Main:    data.Main,
		Weather: data.Weather[0],
	}

	values.Weather.Icon = fmt.Sprintf("http://openweathermap.org/img/wn/%s@2x.png", values.Weather.Icon)

	return values, nil
}

type locationData struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

func (weather OpenWeather) fetchLocationData(cityName string) (locationData, error) {
	logger := logging.Logger{
		Level: weather.Debug,
	}

	request := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", cityName, weather.APIkey)
	logger.Debug(request)
	resp, err := http.Get(request)

	if err != nil {
		logger.Debug(err.Error())
		return locationData{}, errors.New("error fetching data")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Debug(err.Error())
		return locationData{}, errors.New("error reading body")
	}

	var data []locationData

	logger.Debug(string(body))

	err = json.Unmarshal(body, &data)

	if err != nil {
		logger.Debug(err.Error())
		return locationData{}, errors.New("unmarshal error")
	}

	if len(data) == 0 {
		return locationData{}, errors.New("empty response struct")
	}

	return data[0], nil
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
