package service

type Service interface {
	HasPermission(permission string) (bool, error)
}

type service struct {
}

func New() *service {
	return new(service)
}

func (s *service) HasPermission(permission string) (bool, error) {
	return true, nil
}
