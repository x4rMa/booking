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

type FindByToken func(string) (*[]booking_model.Model, error)

type handler struct {
	list        List
	findByToken FindByToken
}

func New(list List, findByToken FindByToken) *handler {
	h := new(handler)
	h.list = list
	h.findByToken = findByToken
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
	if len(request.Form["token"]) > 0 {
		logger.Debugf("token: %s", request.Form["token"][0])
		if list, err = h.findByToken(request.Form["token"][0]); err != nil {
			logger.Debugf("find model by token failed: %v", err)
			return err
		}
	} else {
		if list, err = h.list(); err != nil {
			logger.Debugf("list models failed: %v", err)
			return err
		}
	}
	logger.Debugf("found %d models", len(*list))
	j := json_handler.NewJsonHandler(*list)
	j.ServeHTTP(responseWriter, request)
	return nil
}
