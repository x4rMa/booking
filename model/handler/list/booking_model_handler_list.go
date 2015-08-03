package list

import (
	booking_model "github.com/bborbe/booking/model"
	booking_model_service "github.com/bborbe/booking/model/service"

	"net/http"

	"github.com/bborbe/log"
	error_handler "github.com/bborbe/server/handler/error"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type handler struct {
	service booking_model_service.Service
}

func New(service booking_model_service.Service) *handler {
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
	var list *[]booking_model.Model
	logger.Debug("model list")
	err = request.ParseForm()
	if err != nil {
		return err
	}
	if len(request.Form["token"]) > 0 {
		logger.Debugf("token: %s", request.Form["token"][0])
		if list, err = h.service.FindByToken(request.Form["token"][0]); err != nil {
			logger.Debugf("find model by token failed: %v", err)
			return err
		}
	} else {
		if list, err = h.service.List(); err != nil {
			logger.Debugf("list models failed: %v", err)
			return err
		}
	}
	logger.Debugf("found %d models", len(*list))
	j := json_handler.NewJsonHandler(*list)
	j.ServeHTTP(responseWriter, request)
	return nil
}
