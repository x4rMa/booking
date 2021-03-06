package service

import (
	"fmt"

	booking_model "github.com/bborbe/booking/model"
	booking_model_storage "github.com/bborbe/booking/model/storage"
	"github.com/bborbe/log"
	_ "github.com/lib/pq"
)

var (
	logger = log.DefaultLogger
)

const tokenRetryLimit = 10

type TokenGenerator interface {
	GenerateToken() (string, error)
}

type Service interface {
	List() (*[]booking_model.Model, error)
	Get(id int) (*booking_model.Model, error)
	Create(model *booking_model.Model) (*booking_model.Model, error)
	Delete(id int) (*booking_model.Model, error)
	Update(model *booking_model.Model) (*booking_model.Model, error)
	GetByToken(token string) (*booking_model.Model, error)
	VerifyLogin(model *booking_model.Model) (bool, error)
}

type modelService struct {
	storage        booking_model_storage.Storage
	tokenGenerator TokenGenerator
}

func New(storage booking_model_storage.Storage, tokenGenerator TokenGenerator) *modelService {
	d := new(modelService)
	d.storage = storage
	d.tokenGenerator = tokenGenerator

	return d
}

func (s *modelService) Create(d *booking_model.Model) (*booking_model.Model, error) {
	logger.Debug("create")
	token, err := s.generateToken()
	if err != nil {
		return nil, err
	}
	d.Token = token
	return s.storage.Create(d)
}

func (s *modelService) generateToken() (string, error) {
	for i := 0; i < tokenRetryLimit; i++ {
		token, err := s.tokenGenerator.GenerateToken()
		if err != nil {
			return "", err
		}
		model, err := s.GetByToken(token)
		if err != nil {
			return "", err
		}
		if model == nil {
			return token, nil
		}
	}
	return "", fmt.Errorf("generate token failed")
}

func (s *modelService) VerifyLogin(m *booking_model.Model) (bool, error) {
	model, err := s.GetByToken(m.Token)
	if err != nil {
		return false, err
	}
	return model != nil, nil
}

func (s *modelService) GetByToken(token string) (*booking_model.Model, error) {
	return s.storage.GetByToken(token)
}

func (s *modelService) Update(d *booking_model.Model) (*booking_model.Model, error) {
	logger.Debug("update")
	return s.storage.Update(d)
}

func (s *modelService) List() (*[]booking_model.Model, error) {
	logger.Debug("list")
	return s.storage.Find()
}

func (s *modelService) Get(id int) (*booking_model.Model, error) {
	logger.Debug("get")
	return s.storage.Get(id)
}

func (s *modelService) Delete(id int) (*booking_model.Model, error) {
	logger.Debug("delete")
	return s.storage.Delete(id)
}
