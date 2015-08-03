package list

import (
	booking_shooting "github.com/bborbe/booking/shooting"
	booking_shooting_service "github.com/bborbe/booking/shooting/service"

	"net/http"

	"github.com/bborbe/log"
	error_handler "github.com/bborbe/server/handler/error"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type handler struct {
	service booking_shooting_service.Service
}

func New(service booking_shooting_service.Service) *handler {
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
	var err error
	var list *[]booking_shooting.Shooting
	if list, err = h.service.List(); err != nil {
		logger.Debugf("list shootings failed: %v", err)
		return err
	}
	logger.Debugf("found %d shootings", len(*list))
	j := json_handler.NewJsonHandler(*list)
	j.ServeHTTP(responseWriter, request)
	return nil
}
