package create

import (
	booking_shooting "github.com/bborbe/booking/shooting"
	booking_shooting_service "github.com/bborbe/booking/shooting/service"

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
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	var f booking_shooting.Shooting
	err = json.Unmarshal(content, &f)
	if err != nil {
		return err
	}
	obj, err := h.service.Create(&f)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
