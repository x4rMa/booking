package permission_check_handler

import (
	"fmt"
	"net/http"

	booking_authentication "github.com/bborbe/booking/authentication"
	booking_authorization "github.com/bborbe/booking/authorization"
	booking_handler "github.com/bborbe/booking/handler"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type HasRole func(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error)
type HttpRequestToAuthentication func(request *http.Request) (*booking_authentication.Authentication, error)
type VerifyLogin func(authentication *booking_authentication.Authentication) (bool, error)

type handler struct {
	requiredRole                booking_authorization.Role
	hasRole                     HasRole
	verifyLogin                 VerifyLogin
	httpRequestToAuthentication HttpRequestToAuthentication
	subHandler                  booking_handler.Handler
}

func New(verifyLogin VerifyLogin, hasRole HasRole, httpRequestToAuthentication HttpRequestToAuthentication, requiredRole booking_authorization.Role, subHandler booking_handler.Handler) *handler {
	h := new(handler)
	h.verifyLogin = verifyLogin
	h.requiredRole = requiredRole
	h.hasRole = hasRole
	h.httpRequestToAuthentication = httpRequestToAuthentication
	h.subHandler = subHandler
	return h
}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) error {
	logger.Debug("check permission")
	if err := h.checkPermission(req); err != nil {
		logger.Debugf("check permission failed: %v", err)
		return err
	}
	logger.Debug("permission granted")
	return h.subHandler.ServeHTTP(resp, req)
}

func (h *handler) checkPermission(req *http.Request) error {
	if h.requiredRole == booking_authorization.None {
		return nil
	}
	authentication, err := h.httpRequestToAuthentication(req)
	if err != nil {
		logger.Debugf("httpRequestToAuthentication failed: %v", err)
		return err
	}
	valid, err := h.verifyLogin(authentication)
	if err != nil {
		logger.Debugf("verifyLogin failed: %v", err)
		return err
	}
	if !valid {
		logger.Debugf("validLogin false")
		return fmt.Errorf("permission denied")
	}
	hasRole, err := h.hasRole(authentication, h.requiredRole)
	if err != nil {
		logger.Debugf("hasRole failed: %v", err)
		return err
	}
	if !hasRole {
		logger.Debugf("hasRole false")
		return fmt.Errorf("permission denied")
	}
	return nil
}
