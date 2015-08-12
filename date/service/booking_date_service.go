package service

import (
	booking_date "github.com/bborbe/booking/date"
	booking_date_storage "github.com/bborbe/booking/date/storage"
	"github.com/bborbe/log"
	_ "github.com/lib/pq"
)

var (
	logger = log.DefaultLogger
)

type Service interface {
	List() (*[]booking_date.Date, error)
	ListFree() (*[]booking_date.Date, error)
	Get(id int) (*booking_date.Date, error)
	Create(date *booking_date.Date) (*booking_date.Date, error)
	Delete(id int) (*booking_date.Date, error)
	Update(d *booking_date.Date) (*booking_date.Date, error)
}

type dateService struct {
	storage booking_date_storage.Storage
}

func New(storage booking_date_storage.Storage) *dateService {
	d := new(dateService)
	d.storage = storage
	return d
}

func (s *dateService) Create(d *booking_date.Date) (*booking_date.Date, error) {
	logger.Debug("create date")
	return s.storage.Create(d)
}

func (s *dateService) Update(d *booking_date.Date) (*booking_date.Date, error) {
	logger.Debug("update date")
	return s.storage.Update(d)
}

func (s *dateService) List() (*[]booking_date.Date, error) {
	logger.Debug("list dates")
	return s.storage.Find()
}

func (s *dateService) ListFree() (*[]booking_date.Date, error) {
	logger.Debug("list free dates")
	return s.storage.FindWithoutShootingAndInFuture()
}

func (s *dateService) Get(id int) (*booking_date.Date, error) {
	logger.Debugf("get date with id %d", id)
	return s.storage.Get(id)
}

func (s *dateService) Delete(id int) (*booking_date.Date, error) {
	logger.Debugf("delete date with id %d", id)
	return s.storage.Delete(id)
}
