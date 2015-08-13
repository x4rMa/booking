package service

import (
	booking_authentication "github.com/bborbe/booking/authentication"
	booking_user "github.com/bborbe/booking/user"
	"github.com/bborbe/log"

	booking_model "github.com/bborbe/booking/model"
	model_service "github.com/bborbe/booking/model/service"
	user_service "github.com/bborbe/booking/user/service"
)

var (
	logger = log.DefaultLogger
)

type Service interface {
	VerifyLogin(authentication *booking_authentication.Authentication) (bool, error)
	VerifyModel(authentication *booking_authentication.Authentication) (bool, error)
	VerifyUser(authentication *booking_authentication.Authentication) (bool, error)
}

type authenticationService struct {
	userService  user_service.Service
	modelService model_service.Service
}

func New(userService user_service.Service, modelService model_service.Service) *authenticationService {
	d := new(authenticationService)
	d.userService = userService
	d.modelService = modelService
	return d
}

func (s *authenticationService) VerifyModel(authentication *booking_authentication.Authentication) (bool, error) {
	if len(authentication.Token) > 0 {
		valid, err := s.modelService.VerifyLogin(&booking_model.Model{Token: authentication.Token})
		if err != nil {
			return false, err
		}
		if valid {
			logger.Debug("found valid model")
			return true, nil
		}
	}
	logger.Debug("found no valid model")
	return false, nil
}

func (s *authenticationService) VerifyUser(authentication *booking_authentication.Authentication) (bool, error) {
	logger.Debugf("verify login for authentication: %s", authentication.Login)
	if len(authentication.Login) > 0 && len(authentication.Password) > 0 {
		valid, err := s.userService.VerifyLogin(&booking_user.User{Login: authentication.Login, Password: authentication.Password})
		if err != nil {
			return false, err
		}
		if valid {
			logger.Debug("found valid user")
			return true, nil
		}
	}
	logger.Debug("found no valid user")
	return false, nil
}

func (s *authenticationService) VerifyLogin(authentication *booking_authentication.Authentication) (bool, error) {
	if valid, err := s.VerifyUser(authentication); valid || err != nil {
		return valid, err
	}
	if valid, err := s.VerifyModel(authentication); valid || err != nil {
		return valid, err
	}
	return false, nil
}
