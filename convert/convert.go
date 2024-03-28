package convert

import (
	"github.com/oyevamos/weather_tracker.git/domain"
	"github.com/oyevamos/weather_tracker.git/models"
	"time"
)

func WeatherDataToDomain(data models.WeatherData) domain.Weather {
	now := time.Now()

	return domain.Weather{
		City:    data.Name,
		Celsius: data.Main.Celsius,
		/// comment
		Kelvin:  data.Main.Celsius + 273.15,
		Date:    now,
	}
}
