package error_handler

import (
	"net/http"

	"github.com/bborbe/log"
	error_handler "github.com/bborbe/server/handler/error"
)

var (
	logger = log.DefaultLogger
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request) error
}

type handler struct {
	handler Handler
}

func New(ha Handler) *handler {
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
