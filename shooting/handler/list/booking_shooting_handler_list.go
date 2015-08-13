package list

import (
	"net/http"

	booking_authentication "github.com/bborbe/booking/authentication"
	booking_model "github.com/bborbe/booking/model"
	booking_shooting "github.com/bborbe/booking/shooting"
	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type List func() (*[]booking_shooting.Shooting, error)
type IsParticipant func(authentication *booking_authentication.Authentication) bool
type HttpRequestToAuthentication func(request *http.Request) (*booking_authentication.Authentication, error)
type GetByToken func(token string) (*booking_model.Model, error)
type FindByModelId func(modelId int) (*[]booking_shooting.Shooting, error)

type handler struct {
	list                        List
	isParticipant               IsParticipant
	httpRequestToAuthentication HttpRequestToAuthentication
	getByToken                  GetByToken
	findByModelId               FindByModelId
}

func New(list List, isParticipant IsParticipant, httpRequestToAuthentication HttpRequestToAuthentication, getByToken GetByToken, findByModelId FindByModelId) *handler {
	h := new(handler)
	h.list = list
	h.isParticipant = isParticipant
	h.httpRequestToAuthentication = httpRequestToAuthentication
	h.getByToken = getByToken
	h.findByModelId = findByModelId
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	authentication, err := h.httpRequestToAuthentication(request)
	if err != nil {
		return err
	}
	var list *[]booking_shooting.Shooting
	if h.isParticipant(authentication) {
		model, err := h.getByToken(authentication.Token)
		if err != nil {
			return err
		}
		if list, err = h.findByModelId(model.Id); err != nil {
			logger.Debugf("list shootings failed: %v", err)
			return err
		}
	} else {
		if list, err = h.list(); err != nil {
			logger.Debugf("list shootings failed: %v", err)
			return err
		}
	}
	logger.Debugf("found %d shootings", len(*list))
	j := json_handler.NewJsonHandler(*list)
	j.ServeHTTP(responseWriter, request)
	return nil
}
