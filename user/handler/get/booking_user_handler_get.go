package get

import (
	"net/http"

	"github.com/bborbe/booking/user"
	"github.com/bborbe/log"
	error_handler "github.com/bborbe/server/handler/error"
	json_handler "github.com/bborbe/server/handler/json"
	"github.com/bborbe/server/idparser"
)

var (
	logger = log.DefaultLogger
)

type Get func(int) (*user.User, error)

type handler struct {
	get Get
}

func New(get Get) http.Handler {
	h := new(handler)
	h.get = get
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
	value, err := idparser.ParseIdFormRequest(request)
	if err != nil {
		return err
	}
	obj, err := h.get(value)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
