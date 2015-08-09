package get

import (
	"net/http"

	shooting "github.com/bborbe/booking/shooting"
	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
	"github.com/bborbe/server/idparser"
)

var (
	logger = log.DefaultLogger
)

type Get func(int) (*shooting.Shooting, error)

type handler struct {
	get Get
}

func New(get Get) *handler {
	h := new(handler)
	h.get = get
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
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
