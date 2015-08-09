package get

import (
	"net/http"

	booking_date "github.com/bborbe/booking/date"
	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
	"github.com/bborbe/server/idparser"
)

var (
	logger = log.DefaultLogger
)

type Get func(int) (*booking_date.Date, error)

type handler struct {
	get Get
}

func New(get Get) *handler {
	h := new(handler)
	h.get = get
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	id, err := idparser.ParseIdFormRequest(request)
	if err != nil {
		return err
	}
	obj, err := h.get(id)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
