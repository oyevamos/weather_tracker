package main

import (
	"context"
	"encoding/json"
	"github.com/oyevamos/weather_tracker.git/models"
	"io"
	"log"
	"net/http"
	"time"
)

type weatherController struct {
	service WeatherService
}

func newWeatherController(service WeatherService) weatherController {
	return weatherController{
		service: service,
	}
}

func (wc weatherController) helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go!\n"))
}

func (wc weatherController) weatherHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	switch r.Method {
	case http.MethodDelete:
		var req models.DeleteWeatherRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = wc.service.deleteWeather(ctx, date)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodPost:
		err = wc.service.addWeather(ctx)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func setupWeatherRoutes(controller weatherController) {
	http.HandleFunc("/weather", controller.weatherHandler)
}
