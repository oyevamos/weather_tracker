package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

// import (
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/hello", helloHandler)
// 	setupWeatherRoutes()
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	apiConfig, err := loadApiConfig("C:/Users/User/Desktop/FILES/go/weather_tracker/weather_tracker/.apiConfig")
	if err != nil {
		log.Fatal("Ошибка при загрузке конфигурации: ", err)
	}
	port := apiConfig.Port
	if port == "" {
		port = "0"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
	defer listener.Close()

	fmt.Printf("Сервер запущен на порту %d\n", listener.Addr().(*net.TCPAddr).Port)

	setupWeatherRoutes()
	http.HandleFunc("/hello", helloHandler)

	http.Serve(listener, nil)
}
