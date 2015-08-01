package service

import (
	"github.com/bborbe/booking/date"
	"github.com/bborbe/booking/date/storage"
	"github.com/bborbe/log"
	_ "github.com/lib/pq"
)

var logger = log.DefaultLogger

type Service interface {
	List() (*[]date.Date, error)
	Create(date *date.Date) (*date.Date, error)
}

type dateService struct {
	storage storage.Storage
}

func New(storage storage.Storage) *dateService {
	d := new(dateService)
	d.storage = storage
	return d
}

func (d *dateService) List() (*[]date.Date, error) {
	logger.Debug("List")
	return d.storage.FindDates()
}

func (d *dateService) Create(date *date.Date) (*date.Date, error) {
	return nil, nil
}
