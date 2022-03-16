package main

import (
	"app/weather"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var running = true
	reader := bufio.NewReader(os.Stdin)

	weather.EnableDebug(false)
	weather.SetAPIKey("2bcaa7c6f614a571b27b5e57126aaa30") // enter openweathermap api key

	for running {
		fmt.Print("Enter city name: ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\r\n", "", -1)

		if text == "exit" {
			running = false
			return
		}

		weather.UseCelsius(true)
		response, locationData, err := weather.GetWeather(text)

		if err != nil {
			fmt.Println("Encountered errors while looking up that city, maybe you've entered invalid city?")
			continue
		}

		fmt.Printf("City: %s, Temperature: %f, Feels like: %f, Pressure: %f, Humidity: %f\n",
			locationData.Name, response.Main.Temp, response.Main.FeelsLike, response.Main.Pressure, response.Main.Humidity)
		fmt.Println("\nEnter 'exit' to exit next time")
	}
}
