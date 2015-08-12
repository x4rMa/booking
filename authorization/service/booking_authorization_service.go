package service

import (
	booking_authentication "github.com/bborbe/booking/authentication"
	booking_authorization "github.com/bborbe/booking/authorization"
)

type VerifyLogin func(authentication *booking_authentication.Authentication) (bool, error)

type Service interface {
	HasRole(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error)
}

type service struct {
	verifyLogin VerifyLogin
}

func New(verifyLogin VerifyLogin) *service {
	s := new(service)
	s.verifyLogin = verifyLogin
	return s
}

func (s *service) HasRole(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error) {
	if booking_authorization.None == role {
		return true, nil
	}
	valid, err := s.verifyLogin(authentication)
	if err != nil || !valid {
		return valid, err
	}
	if booking_authorization.Administrator == role && len(authentication.Token) == 0 && authentication.Login == "admin" {
		return true, nil
	}
	if booking_authorization.Organizer == role && len(authentication.Token) == 0 && len(authentication.Login) > 0 {
		return true, nil
	}
	if booking_authorization.Participant == role && (len(authentication.Token) > 0 && len(authentication.Login) == 0 || len(authentication.Token) == 0 && len(authentication.Login) > 0) {
		return true, nil
	}
	return false, nil
}
