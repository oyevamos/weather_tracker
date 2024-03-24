package models

import "time"

type GetWeatherRequest struct {
	City string    `json:"city"`
	Date time.Time `json:"date"`
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

type DeleteWeatherRequest struct {
	Date string `json:"date"`
}
