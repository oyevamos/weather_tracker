package config

import (
	"encoding/json"
	"io/ioutil"
)

type Postgres struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type ApiConfigData struct {
	OpenWeatherMapApiKey string   `json:"OpenWeatherMapApiKey"`
	City                 string   `json:"City"`
	Port                 string   `json:"Port"`
	Postgres             Postgres `json:"Postgres"`
}

func LoadApiConfig(filename string) (ApiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ApiConfigData{}, err
	}
	var c ApiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return ApiConfigData{}, err
	}
	return c, nil
}
