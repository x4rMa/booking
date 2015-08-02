package service

import (
	"github.com/bborbe/booking/date"
	"github.com/bborbe/booking/date/storage"
	_ "github.com/lib/pq"
)

type Service interface {
	List() (*[]date.Date, error)
	Get(id int) (*date.Date, error)
	Create(date *date.Date) (*date.Date, error)
	Delete(id int) (*date.Date, error)
	Update(d *date.Date) (*date.Date, error)
}

type dateService struct {
	storage storage.Storage
}

func New(storage storage.Storage) *dateService {
	d := new(dateService)
	d.storage = storage
	return d
}

func (s *dateService) Create(d *date.Date) (*date.Date, error) {
	return s.storage.Create(d)
}

func (s *dateService) Update(d *date.Date) (*date.Date, error) {
	return s.storage.Update(d)
}

func (s *dateService) List() (*[]date.Date, error) {
	return s.storage.Find()
}

func (s *dateService) Get(id int) (*date.Date, error) {
	return s.storage.Get(id)
}

func (s *dateService) Delete(id int) (*date.Date, error) {
	return s.storage.Delete(id)
}
