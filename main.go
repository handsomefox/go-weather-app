package main

import (
	"app/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

func main() {
	app.Route("/", &components.WeatherInformation{})
	app.RunWhenOnBrowser()

	app.Link().Rel("preconnect").Href("https://fonts.gstatic.com")
	app.Link().Rel("stylesheet").Href("https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,400;1,300&display=swap")

	http.Handle("/", &app.Handler{
		Name:        "MyWeatherApp",
		Description: "Weather app using go!",
		Styles: []string{
			"/web/style.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
