package domain

import "time"

type Weather struct {
	Kelvin  float64
	Celsius float64
	City    string
	Date    time.Time
}
