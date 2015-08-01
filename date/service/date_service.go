package service

import "github.com/bborbe/booking/date"

type DateService interface {
	List() []date.Date
	Create(date date.Date) error
}

type dateService struct {

}

func New() *dateService {
	return new(dateService)
}

func (d *dateService) List() []date.Date {
	return make([]date.Date, 0)
}

func (d *dateService) Create(date date.Date) error {
	return nil
}
