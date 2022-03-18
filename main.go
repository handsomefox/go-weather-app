package main

import (
	"app/logging"
	"app/openweather"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var running = true
	reader := bufio.NewReader(os.Stdin)

	weather := openweather.Weather{
		APIkey:  "2bcaa7c6f614a571b27b5e57126aaa30", // enter openweathermap api key
		Celsius: true,
		Debug:   logging.OFF,
	}

	for running {
		fmt.Print("Enter city name: ")
		text, _ := reader.ReadString('\n')

		start := time.Now()

		text = strings.Replace(text, "\r\n", "", -1)
		if text == "exit" {
			running = false
			return
		}

		data, err := weather.Get(text)
		if err != nil {
			fmt.Println("Encountered errors while looking up that city, maybe you've entered invalid city?")
			continue
		}

		fmt.Printf("City: %s, Temperature: %.2f, Feels like: %.2f, Pressure: %.0f, Humidity: %.0f\n",
			data.Name, data.Main.Temp, data.Main.FeelsLike, data.Main.Pressure, data.Main.Humidity)

		elapsed := time.Since(start)

		fmt.Printf("Elapsed: %dms\n", elapsed.Milliseconds())
		fmt.Println("\nEnter 'exit' to exit next time")
	}
}
