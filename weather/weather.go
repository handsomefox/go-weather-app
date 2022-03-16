package weather

import (
	"app/logging"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiKey string

var useCelsius = false
var enableDebug = false

func SetAPIKey(key string){
	apiKey = key
}

func UseCelsius(state bool) {
	useCelsius = state
}

func EnableDebug(state bool) {
	enableDebug = state
}

func GetWeather(city string) (Response, LocationData, error) {
	logging.SetLogLevel(logging.DEBUG)
	locationData, err := fetchLocationData(city)

	if err != nil {
		if enableDebug {
			logging.Debug(err.Error())
		}
		return Response{}, LocationData{}, errors.New("error fetching location data")
	}

	request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s",
		locationData.Lat, locationData.Lon, apiKey)

	resp, err := http.Get(request)

	if err != nil {
		if enableDebug {
			logging.Debug(err.Error())
		}
		return Response{}, LocationData{}, errors.New("error fetching weather data")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		if enableDebug {
			logging.Debug(err.Error())
		}
		return Response{}, LocationData{}, errors.New("error reading response body")
	}

	var data Response
	err = json.Unmarshal(body, &data)

	if err != nil {
		if enableDebug {
			logging.Debug(err.Error())
		}
		return Response{}, LocationData{}, errors.New("unmarshal error")
	}

	if enableDebug {
		fmt.Println(data)
	}

	if useCelsius {
		data.Main.Temp = kelvinToCelsius(data.Main.Temp)
		data.Main.TempMax = kelvinToCelsius(data.Main.TempMax)
		data.Main.TempMin = kelvinToCelsius(data.Main.TempMin)
		data.Main.FeelsLike = kelvinToCelsius(data.Main.FeelsLike)
	}

	return data, locationData, nil
}

func fetchLocationData(cityName string) (LocationData, error) {
	logging.SetLogLevel(logging.DEBUG)
	request := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", cityName, apiKey)
	resp, err := http.Get(request)

	if err != nil {
		if enableDebug {
			logging.Debug(err.Error())
		}
		return LocationData{}, errors.New("error fetching data")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		if enableDebug {
			logging.Debug(err.Error())
		}
		return LocationData{}, errors.New("error reading body")
	}

	var data []LocationData

	if enableDebug {
		logging.Debug(string(body))
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		if enableDebug {
			logging.Debug(err.Error())
		}
		return LocationData{}, errors.New("unmarshal error")
	}
	if enableDebug {
		fmt.Println(data)
	}

	if len(data) == 0 {
		return LocationData{}, errors.New("empty response struct")
	}

	return data[0], nil
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
