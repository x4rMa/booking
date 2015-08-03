package service

import (
	"github.com/bborbe/booking/shooting"
	"github.com/bborbe/booking/shooting/storage"
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
}

type shootingService struct {
	storage storage.Storage
}

func New(storage storage.Storage) *shootingService {
	d := new(shootingService)
	d.storage = storage
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
