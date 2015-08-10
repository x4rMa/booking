package permission_check_handler

import (
	"fmt"
	"net/http"

	booking_authentication "github.com/bborbe/booking/authentication"
	booking_authorization "github.com/bborbe/booking/authorization"
	booking_handler "github.com/bborbe/booking/handler"
)

type HasRole func(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error)
type HttpRequestToAuthentication func(request *http.Request) (*booking_authentication.Authentication, error)

type handler struct {
	requiredRole                booking_authorization.Role
	hasRole                     HasRole
	httpRequestToAuthentication HttpRequestToAuthentication
	subHandler                  booking_handler.Handler
}

func New(hasRole HasRole, httpRequestToAuthentication HttpRequestToAuthentication, requiredRole booking_authorization.Role, subHandler booking_handler.Handler) *handler {
	h := new(handler)
	h.requiredRole = requiredRole
	h.hasRole = hasRole
	h.httpRequestToAuthentication = httpRequestToAuthentication
	h.subHandler = subHandler
	return h
}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) error {
	if err := h.checkPermission(req); err != nil {
		return err
	}
	return h.subHandler.ServeHTTP(resp, req)
}

func (h *handler) checkPermission(req *http.Request) error {
	if h.requiredRole == booking_authorization.None {
		return nil
	}
	authentication, err := h.httpRequestToAuthentication(req)
	if err != nil {
		return err
	}
	hasRole, err := h.hasRole(authentication, h.requiredRole)
	if err != nil {
		return err
	}
	if !hasRole {
		return fmt.Errorf("permission denied")
	}
	return nil
}
