package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	setupWeatherRoutes()
	http.ListenAndServe(":8080", nil)
}
