package service

import (
	booking_authentication "github.com/bborbe/booking/authentication"
	booking_authorization "github.com/bborbe/booking/authorization"
)

type Service interface {
	HasRole(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error)
}

type service struct {
}

func New() *service {
	return new(service)
}

func (s *service) HasRole(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error) {
	return true, nil
}
