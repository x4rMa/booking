package service

import (
	"fmt"

	"github.com/bborbe/booking/model"
	"github.com/bborbe/booking/model/storage"
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
	List() (*[]model.Model, error)
	Get(id int) (*model.Model, error)
	Create(model *model.Model) (*model.Model, error)
	Delete(id int) (*model.Model, error)
	Update(model *model.Model) (*model.Model, error)
	FindByToken(token string) (*[]model.Model, error)
	VerifyLogin(model *model.Model) (bool, error )
}

type modelService struct {
	storage        storage.Storage
	tokenGenerator TokenGenerator
}

func New(storage storage.Storage, tokenGenerator TokenGenerator) *modelService {
	d := new(modelService)
	d.storage = storage
	d.tokenGenerator = tokenGenerator

	return d
}

func (s *modelService) Create(d *model.Model) (*model.Model, error) {
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
		models, err := s.FindByToken(token)
		if err != nil {
			return "", err
		}
		if len(*models) == 0 {
			return token, nil
		}
	}
	return "", fmt.Errorf("generate token failed")
}

func (s *modelService) VerifyLogin(m *model.Model) (bool, error ) {
	list, err := s.FindByToken(m.Token)
	if err != nil {
		return false, err
	}
	return len(*list) > 0, nil
}

func (s *modelService) FindByToken(token string) (*[]model.Model, error) {
	return s.storage.FindByToken(token)
}

func (s *modelService) Update(d *model.Model) (*model.Model, error) {
	logger.Debug("update")
	return s.storage.Update(d)
}

func (s *modelService) List() (*[]model.Model, error) {
	logger.Debug("list")
	return s.storage.Find()
}

func (s *modelService) Get(id int) (*model.Model, error) {
	logger.Debug("get")
	return s.storage.Get(id)
}

func (s *modelService) Delete(id int) (*model.Model, error) {
	logger.Debug("delete")
	return s.storage.Delete(id)
}
