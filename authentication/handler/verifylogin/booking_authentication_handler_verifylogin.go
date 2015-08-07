package verifylogin

import (
	booking_authentication "github.com/bborbe/booking/authentication"
	booking_authentication_service "github.com/bborbe/booking/authentication/service"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bborbe/log"
	error_handler "github.com/bborbe/server/handler/error"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type handler struct {
	service booking_authentication_service.Service
}

func New(service booking_authentication_service.Service) *handler {
	h := new(handler)
	h.service = service
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	err := h.serveHTTP(responseWriter, request)
	if err != nil {
		logger.Debug(err)
		e := error_handler.NewErrorMessage(http.StatusInternalServerError, err.Error())
		e.ServeHTTP(responseWriter, request)
		return
	}
}

func (h *handler) serveHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	logger.Debugf("verifylogin authentication: %v", string(content))
	var f booking_authentication.Authentication
	err = json.Unmarshal(content, &f)
	if err != nil {
		return err
	}
	obj, err := h.service.VerifyLogin(&f)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}