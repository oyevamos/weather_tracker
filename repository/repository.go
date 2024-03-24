package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/oyevamos/weather_tracker.git/config"
	"github.com/oyevamos/weather_tracker.git/domain"
	"time"
)

type WeatherRepository struct {
	db *sql.DB
}

func NewWeather(cfg config.Postgres) (*WeatherRepository, error) {
	db, err := connectDB(cfg)
	if err != nil {
		return nil, err
	}
	return &WeatherRepository{
		db: db,
	}, nil
}
func connectDB(cfg config.Postgres) (*sql.DB, error) {
	connectURL := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)

	db, err := sql.Open("postgres", connectURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (w *WeatherRepository) AddWeather(ctx context.Context, weather domain.Weather) error {
	query := `
	INSERT INTO weather (kelvin, celsius, city, date) VALUES($1, $2, $3, $4);
`
	_, err := w.db.ExecContext(ctx, query, weather.Kelvin, weather.Celsius, weather.City, weather.Date)
	return err
}

func (w *WeatherRepository) DeleteWeather(ctx context.Context, date time.Time) error {
	query := `
	DELETE FROM weather WHERE date = $1;
`
	_, err := w.db.ExecContext(ctx, query, date)
	return err
}
