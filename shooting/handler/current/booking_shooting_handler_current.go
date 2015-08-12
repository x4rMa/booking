package current

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

type FindByModelId func(modelId int) (*[]booking_shooting.Shooting, error)
type HttpRequestToAuthentication func(request *http.Request) (*booking_authentication.Authentication, error)
type GetByToken func(token string) (*booking_model.Model, error)

type handler struct {
	httpRequestToAuthentication HttpRequestToAuthentication
	getByToken                  GetByToken
	findByModelId               FindByModelId
}

func New(httpRequestToAuthentication HttpRequestToAuthentication, getByToken GetByToken, findByModelId FindByModelId) *handler {
	h := new(handler)
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
	model, err := h.getByToken(authentication.Token)
	if err != nil {
		return err
	}
	var list *[]booking_shooting.Shooting
	if list, err = h.findByModelId(model.Id); err != nil {
		logger.Debugf("list shootings failed: %v", err)
		return err
	}
	logger.Debugf("found %d shootings", len(*list))
	j := json_handler.NewJsonHandler(*list)
	j.ServeHTTP(responseWriter, request)
	return nil
}
