package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/oyevamos/weather_tracker.git/config"
	"log"
)

type Weather struct {
	db *sql.DB
}

func NewWeather(cfg config.Postgres) (*Weather, error) {
	db, err := connectDB(cfg)
	if err != nil {
		return nil, err
	}
	return &Weather{
		db: db,
	}, nil
}
func connectDB(cfg config.Postgres) (*sql.DB, error) {
	log.Println(cfg)
	connectURL := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)
	log.Println(connectURL)
	db, err := sql.Open("postgres", connectURL)
	if err != nil {
		return nil, err
	}

	log.Println(1)
	err = db.Ping()

	if err != nil {
		return nil, err
	}
	log.Println(2)
	return db, nil
}
