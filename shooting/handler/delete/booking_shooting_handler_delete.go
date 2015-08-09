package delete

import (
	"net/http"

	booking_shooting "github.com/bborbe/booking/shooting"
	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
	"github.com/bborbe/server/idparser"
)

var (
	logger = log.DefaultLogger
)

type Delete func(int) (*booking_shooting.Shooting, error)

type handler struct {
	delete Delete
}

func New(delete Delete) *handler {
	h := new(handler)
	h.delete = delete
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	id, err := idparser.ParseIdFormRequest(request)
	if err != nil {
		return err
	}
	obj, err := h.delete(id)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
