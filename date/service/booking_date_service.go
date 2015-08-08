package service

import (
	"github.com/bborbe/booking/date"
	"github.com/bborbe/booking/date/storage"
	"github.com/bborbe/log"
	_ "github.com/lib/pq"
)

var (
	logger = log.DefaultLogger
)

type Service interface {
	List() (*[]date.Date, error)
	ListFree() (*[]date.Date, error)
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
	logger.Debug("create date")
	return s.storage.Create(d)
}

func (s *dateService) Update(d *date.Date) (*date.Date, error) {
	logger.Debug("update date")
	return s.storage.Update(d)
}

func (s *dateService) List() (*[]date.Date, error) {
	logger.Debug("list dates")
	return s.storage.Find()
}

func (s *dateService) ListFree() (*[]date.Date, error) {
	logger.Debug("list free dates")
	return s.storage.FindWithoutShooting()
}

func (s *dateService) Get(id int) (*date.Date, error) {
	logger.Debugf("get date with id %d", id)
	return s.storage.Get(id)
}

func (s *dateService) Delete(id int) (*date.Date, error) {
	logger.Debugf("delete date with id %d", id)
	return s.storage.Delete(id)
}
