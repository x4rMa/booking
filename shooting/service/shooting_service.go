package service

import "github.com/bborbe/booking/shooting"

type ShootingService interface {
	List() []shooting.Shooting
	Create(shooting shooting.Shooting) error
}

type shootingService struct {
}

func New() *shootingService {
	return new(shootingService)
}

func (d *shootingService) List() []shooting.Shooting {
	return make([]shooting.Shooting, 0)
}

func (d *shootingService) Create(shooting shooting.Shooting) error {
	return nil
}
