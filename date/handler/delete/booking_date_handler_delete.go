package delete

import (
	"net/http"

	date_service "github.com/bborbe/booking/date/service"
	"github.com/bborbe/log"
	error_handler "github.com/bborbe/server/handler/error"
	json_handler "github.com/bborbe/server/handler/json"
	"github.com/bborbe/server/idparser"
)

var (
	logger = log.DefaultLogger
)

type handler struct {
	dateService date_service.Service
}

func New(dateService date_service.Service) http.Handler {
	h := new(handler)
	h.dateService = dateService
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
	id, err := idparser.ParseIdFormRequest(request)
	if err != nil {
		return err
	}
	obj, err := h.dateService.Delete(id)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
