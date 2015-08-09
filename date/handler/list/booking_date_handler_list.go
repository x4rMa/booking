package list

import (
	"net/http"

	booking_date "github.com/bborbe/booking/date"
	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type List func() (*[]booking_date.Date, error)

type handler struct {
	list List
}

func New(list List) *handler {
	h := new(handler)
	h.list = list
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	var err error
	var list *[]booking_date.Date
	if list, err = h.list(); err != nil {
		logger.Debugf("list dates failed: %v", err)
		return err
	}
	logger.Debugf("found %d dates", len(*list))
	j := json_handler.NewJsonHandler(*list)
	j.ServeHTTP(responseWriter, request)
	return nil
}
