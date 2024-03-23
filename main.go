package main

import (
	"fmt"
	"github.com/oyevamos/weather_tracker.git/config"
	"github.com/oyevamos/weather_tracker.git/repository"
	"log"
	"net"
	"net/http"
)

func main() {
	apiConfig, err := config.LoadApiConfig("C:/Users/User/Desktop/FILES/go/weather_tracker/weather_tracker/.apiConfig")
	if err != nil {
		log.Fatal("Ошибка при загрузке конфигурации: ", err)
	}

	port := apiConfig.Port
	if port == "" {
		port = "0" // показывает что я не хочу конерктизироваить порт, его выберет система.
	}

	_, err = repository.NewWeather(apiConfig.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
	defer listener.Close()

	fmt.Printf("Сервер запущен на порту %d\n", listener.Addr().(*net.TCPAddr).Port) //заклинание для нахождения свободного порта, взято отсюда https://stackoverflow.com/questions/43424787/how-to-use-next-available-port-in-http-listenandserve

	setupWeatherRoutes()
	http.HandleFunc("/hello", helloHandler)

	http.Serve(listener, nil) // Запускаем сервер с использованием созданного listener
}
