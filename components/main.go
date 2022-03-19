package components

import (
	"app/logging"
	"app/openweather"
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type WeatherInformation struct {
	app.Compo
	Input string
	Data  openweather.Values
}

func (c *WeatherInformation) Render() app.UI {

	var src string

	if len(c.Data.Weather.Icon) > 0 {
		src = c.Data.Weather.Icon
	} else {
		src = "http://openweathermap.org/img/wn/01d@2x.png"
	}

	return app.Div().Body(
		app.Div().Class("header").Body(
			app.A().Text("MyWeatherApp"),
		),
		app.Div().Class("input-field").Body(
			app.Label().For("input-label").Text("Enter a city name: ").Class("label"),
			app.Input().Spellcheck(false).Type("text").Class("input-city").ID("location-city").OnChange(c.inputOnChange),
		),
		app.Div().Class("weather-information").Body(
			app.P().Class("weather-temperature").ID("temp").Body(
				app.Text(fmt.Sprintf("Temperature: %.2f°C", c.Data.Main.Temp)),
				app.P().Body(app.Img().Class("img").Src(src)),
			),
		),
	)
}

func (c *WeatherInformation) inputOnChange(ctx app.Context, e app.Event) {
	v := ctx.JSSrc().Get("value")

	weather := openweather.OpenWeather{
		APIkey:  "2bcaa7c6f614a571b27b5e57126aaa30", // enter openweathermap api key
		Celsius: true,
		Debug:   logging.OFF,
	}

	data, err := weather.Get(v.String())
	if err != nil {
		fmt.Println("Encountered errors while looking up that city, maybe you've entered invalid city?")
	}
	c.Data = data

	ctx.SetState("weather-data", &data)
}

func (c *WeatherInformation) OnMount(ctx app.Context) {
	ctx.ObserveState("weather-data").Value(&c.Data)
	ctx.ObserveState("weather-input").Value(&c.Input)
	fmt.Println(c.Data)
}
