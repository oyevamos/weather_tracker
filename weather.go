package main

import (
	"context"
	"encoding/json"
	"github.com/oyevamos/weather_tracker.git/config"
	"github.com/oyevamos/weather_tracker.git/convert"
	"github.com/oyevamos/weather_tracker.git/models"
	"github.com/oyevamos/weather_tracker.git/repository"
	"net/http"
	"time"
)

type WeatherService struct {
	repo *repository.WeatherRepository
}

func NewWeatherService(repo *repository.WeatherRepository) WeatherService {
	return WeatherService{
		repo: repo,
	}
}

func (w WeatherService) addWeather(ctx context.Context) error {
	apiConfig, err := config.LoadApiConfig(".apiConfig")
	if err != nil {
		return err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + apiConfig.City + "&units=metric")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var d models.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return err
	}

	return w.repo.AddWeather(ctx, convert.WeatherDataToDomain(d))
}

func (w WeatherService) deleteWeather(ctx context.Context, date time.Time) error {
	return w.repo.DeleteWeather(ctx, date)
}
