package service

import (
	"database/sql"
	"fmt"

	"github.com/bborbe/booking/date"
	"github.com/bborbe/log"
	_ "github.com/lib/pq"
)

var logger = log.DefaultLogger

type DatabaseConfig interface {
	Username() string
	Password() string
	Hostname() string
	Database() string
}

type DateService interface {
	List() ([]date.Date, error)
	Create(date date.Date) error
}

type dateService struct {
	config DatabaseConfig
}

func New(config DatabaseConfig) *dateService {
	d := new(dateService)
	d.config = config
	return d
}

func (d *dateService) List() ([]date.Date, error) {
	logger.Debug("List")
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://{}:{}@{}/{}?sslmode=verify-full", d.config.Username(), d.config.Password(), d.config.Hostname(), d.config.Database()))
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM date")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Columns()
	}

	return make([]date.Date, 0), nil
}

func (d *dateService) Create(date date.Date) error {
	return nil
}
