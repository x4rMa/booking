package list

import (
	booking_model "github.com/bborbe/booking/model"

	"net/http"

	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type List func() (*[]booking_model.Model, error)


type handler struct {
	list        List
}

func New(list List) *handler {
	h := new(handler)
	h.list = list
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	var err error
	var list *[]booking_model.Model
	logger.Debug("model list")
	err = request.ParseForm()
	if err != nil {
		return err
	}
	if list, err = h.list(); err != nil {
		logger.Debugf("list models failed: %v", err)
		return err
	}
	logger.Debugf("found %d models", len(*list))
	j := json_handler.NewJsonHandler(*list)
	j.ServeHTTP(responseWriter, request)
	return nil
}
