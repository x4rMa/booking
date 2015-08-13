package service

import (
	booking_authentication "github.com/bborbe/booking/authentication"
	booking_authorization "github.com/bborbe/booking/authorization"
)

type Service interface {
	HasRole(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error)
	IsAdministrator(authentication *booking_authentication.Authentication) bool
	IsOrganizer(authentication *booking_authentication.Authentication) bool
	IsParticipant(authentication *booking_authentication.Authentication) bool
}

type service struct {
}

func New() *service {
	s := new(service)
	return s
}

func (s *service) HasRole(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error) {
	if booking_authorization.None == role {
		return true, nil
	}
	if booking_authorization.Administrator == role && s.IsAdministrator(authentication) {
		return true, nil
	}
	if booking_authorization.Organizer == role && (s.IsAdministrator(authentication) || s.IsOrganizer(authentication)) {
		return true, nil
	}
	if booking_authorization.Participant == role && (s.IsAdministrator(authentication) || s.IsOrganizer(authentication) || s.IsParticipant(authentication)) {
		return true, nil
	}
	return false, nil
}

func (s *service) IsAdministrator(authentication *booking_authentication.Authentication) bool {
	return len(authentication.Token) == 0 && authentication.Login == "admin"
}

func (s *service) IsOrganizer(authentication *booking_authentication.Authentication) bool {
	return len(authentication.Token) == 0 && len(authentication.Login) > 0
}

func (s *service) IsParticipant(authentication *booking_authentication.Authentication) bool {
	return len(authentication.Token) > 0 && len(authentication.Login) == 0
}
