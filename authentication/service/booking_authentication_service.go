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
	VerifyLogin(d *booking_authentication.Authentication) (bool, error)
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

func (s *authenticationService) VerifyLogin(auth *booking_authentication.Authentication) (bool, error) {
	logger.Debugf("verify login for authentication: %s", auth.Login)
	if len(auth.Login) > 0 && len(auth.Password) > 0 {
		valid, err := s.userService.VerifyLogin(&booking_user.User{Login: auth.Login, Password: auth.Password})
		if err != nil {
			return false, err
		}
		if valid {
			logger.Debug("found valid user")
			return true, nil
		}
	}
	if len(auth.Token) > 0 {
		valid, err := s.modelService.VerifyLogin(&booking_model.Model{Token: auth.Token})
		if err != nil {
			return false, err
		}
		if valid {
			logger.Debug("found valid model")
			return true, nil
		}
	}
	logger.Debug("found no valid user or model")
	return false, nil
}
