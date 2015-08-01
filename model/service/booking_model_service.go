package service

import "github.com/bborbe/booking/model"

type ModelService interface {
	List() []model.Model
	Create(model model.Model) error
}

type modelService struct {
}

func New() *modelService {
	return new(modelService)
}

func (d *modelService) List() []model.Model {
	return make([]model.Model, 0)
}

func (d *modelService) Create(model model.Model) error {
	return nil
}
