package service

import (
	"github.com/bborbe/booking/authentication"
	"github.com/bborbe/log"
	"github.com/bborbe/booking/user"

	user_service "github.com/bborbe/booking/user/service"
	model_service "github.com/bborbe/booking/model/service"
	"github.com/bborbe/booking/model"
)

var (
	logger = log.DefaultLogger
)

type Service interface {
	VerifyLogin(d *authentication.Authentication) (bool, error)
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

func (s *authenticationService) VerifyLogin(auth *authentication.Authentication) (bool, error) {
	logger.Debugf("verify login for authentication: %s", auth.Login)
	if len(auth.Login) > 0 && len(auth.Password) > 0 {
		valid, err := s.userService.VerifyLogin(&user.User{Login:auth.Login, Password:auth.Password})
		if err != nil {
			return false, err
		}
		if valid {
			logger.Debug("found valid user")
			return true, nil
		}
	}
	if len(auth.Token) > 0 {
		valid, err := s.modelService.VerifyLogin(&model.Model{Token:auth.Token})
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
