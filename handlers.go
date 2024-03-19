package main

import (
	"encoding/json"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go!\n"))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	data, err := queryWeather()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func setupWeatherRoutes() {
	http.HandleFunc("/weather/", weatherHandler)
}
