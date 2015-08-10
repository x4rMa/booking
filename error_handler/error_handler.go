package error_handler

import (
	"net/http"

	booking_handler "github.com/bborbe/booking/handler"
	"github.com/bborbe/log"
	error_handler "github.com/bborbe/server/handler/error"
)

var (
	logger = log.DefaultLogger
)

type handler struct {
	handler booking_handler.Handler
}

func New(ha booking_handler.Handler) *handler {
	h := new(handler)
	h.handler = ha
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	err := h.handler.ServeHTTP(responseWriter, request)
	if err != nil {
		logger.Debug(err)
		e := error_handler.NewErrorMessage(http.StatusInternalServerError, err.Error())
		e.ServeHTTP(responseWriter, request)
		return
	}
}
