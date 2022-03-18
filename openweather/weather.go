package openweather

import (
	"app/logging"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Name string
	Main Main `json:"main"`
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
	APIkey  string
	Celsius bool
	Debug   logging.LogLevel
}

func (weather Weather) Get(city string) (Response, error) {
	logger := logging.Logger{
		Level: weather.Debug,
	}
	locationData, err := weather.fetchLocationData(city)

	if err != nil {
		logger.Debug(err.Error())
		return Response{}, errors.New("error fetching location data")
	}

	request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s",
		locationData.Lat, locationData.Lon, weather.APIkey)

	logger.Debug(request)

	resp, err := http.Get(request)

	if err != nil {
		logger.Debug(err.Error())
		return Response{}, errors.New("error fetching openweather data")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Debug(err.Error())
		return Response{}, errors.New("error reading response body")
	}

	var data Response
	err = json.Unmarshal(body, &data)

	if err != nil {
		logger.Debug(err.Error())
		return Response{}, errors.New("unmarshal error")
	}

	if weather.Celsius {
		data.Main.Temp = kelvinToCelsius(data.Main.Temp)
		data.Main.TempMax = kelvinToCelsius(data.Main.TempMax)
		data.Main.TempMin = kelvinToCelsius(data.Main.TempMin)
		data.Main.FeelsLike = kelvinToCelsius(data.Main.FeelsLike)
	}

	data.Name = locationData.Name
	return data, nil
}

type locationData struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

func (weather Weather) fetchLocationData(cityName string) (locationData, error) {
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
