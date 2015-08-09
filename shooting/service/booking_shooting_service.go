package service

import (
	booking_booked_event "github.com/bborbe/booking/booked_event"
	booking_shooting "github.com/bborbe/booking/shooting"
	booking_shooting_storage "github.com/bborbe/booking/shooting/storage"
	"github.com/bborbe/eventbus"
	"github.com/bborbe/log"
	_ "github.com/lib/pq"
)

var (
	logger = log.DefaultLogger
)

type Service interface {
	List() (*[]booking_shooting.Shooting, error)
	Get(id int) (*booking_shooting.Shooting, error)
	Create(shooting *booking_shooting.Shooting) (*booking_shooting.Shooting, error)
	Delete(id int) (*booking_shooting.Shooting, error)
	Update(d *booking_shooting.Shooting) (*booking_shooting.Shooting, error)
	Book(d *booking_shooting.Shooting) (*booking_shooting.Shooting, error)
}

type shootingService struct {
	storage  booking_shooting_storage.Storage
	eventbus eventbus.EventBus
}

func New(storage booking_shooting_storage.Storage, eventbus eventbus.EventBus) *shootingService {
	d := new(shootingService)
	d.storage = storage
	d.eventbus = eventbus
	return d
}

func (s *shootingService) Create(d *booking_shooting.Shooting) (*booking_shooting.Shooting, error) {
	logger.Debug("create")
	return s.storage.Create(d)
}

func (s *shootingService) Update(d *booking_shooting.Shooting) (*booking_shooting.Shooting, error) {
	logger.Debug("update")
	return s.storage.Update(d)
}

func (s *shootingService) Book(d *booking_shooting.Shooting) (*booking_shooting.Shooting, error) {
	logger.Debug("book")
	obj, err := s.storage.Get(d.Id)
	if err != nil {
		return nil, err
	}
	obj.DateId = d.DateId
	result, err := s.storage.Update(obj)
	if err != nil {
		return nil, err
	}
	s.eventbus.Publish(booking_booked_event.New(*result))
	return result, nil
}

func (s *shootingService) List() (*[]booking_shooting.Shooting, error) {
	logger.Debug("list")
	return s.storage.Find()
}

func (s *shootingService) Get(id int) (*booking_shooting.Shooting, error) {
	logger.Debug("get")
	return s.storage.Get(id)
}

func (s *shootingService) Delete(id int) (*booking_shooting.Shooting, error) {
	logger.Debug("delete")
	return s.storage.Delete(id)
}
