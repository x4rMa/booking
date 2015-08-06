package service

import (
	"github.com/bborbe/booking/booked_event"
	"github.com/bborbe/booking/shooting"
	"github.com/bborbe/booking/shooting/storage"
	"github.com/bborbe/eventbus"
	"github.com/bborbe/log"
	_ "github.com/lib/pq"
)

var (
	logger = log.DefaultLogger
)

type Service interface {
	List() (*[]shooting.Shooting, error)
	Get(id int) (*shooting.Shooting, error)
	Create(shooting *shooting.Shooting) (*shooting.Shooting, error)
	Delete(id int) (*shooting.Shooting, error)
	Update(d *shooting.Shooting) (*shooting.Shooting, error)
	Book(d *shooting.Shooting) (*shooting.Shooting, error)
}

type shootingService struct {
	storage  storage.Storage
	eventbus eventbus.EventBus
}

func New(storage storage.Storage, eventbus eventbus.EventBus) *shootingService {
	d := new(shootingService)
	d.storage = storage
	d.eventbus = eventbus
	return d
}

func (s *shootingService) Create(d *shooting.Shooting) (*shooting.Shooting, error) {
	logger.Debug("create")
	return s.storage.Create(d)
}

func (s *shootingService) Update(d *shooting.Shooting) (*shooting.Shooting, error) {
	logger.Debug("update")
	return s.storage.Update(d)
}

func (s *shootingService) Book(d *shooting.Shooting) (*shooting.Shooting, error) {
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
	s.eventbus.Publish(booked_event.New(*result))
	return result, nil
}

func (s *shootingService) List() (*[]shooting.Shooting, error) {
	logger.Debug("list")
	return s.storage.Find()
}

func (s *shootingService) Get(id int) (*shooting.Shooting, error) {
	logger.Debug("get")
	return s.storage.Get(id)
}

func (s *shootingService) Delete(id int) (*shooting.Shooting, error) {
	logger.Debug("delete")
	return s.storage.Delete(id)
}
