package create

import (
	booking_model "github.com/bborbe/booking/model"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type Create func(*booking_model.Model) (*booking_model.Model, error)

type handler struct {
	create Create
}

func New(create Create) *handler {
	h := new(handler)
	h.create = create
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	logger.Debugf("model create: %s", string(content))
	var f booking_model.Model
	err = json.Unmarshal(content, &f)
	if err != nil {
		return err
	}
	obj, err := h.create(&f)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
