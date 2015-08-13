package get

import (
	"net/http"

	booking_authentication "github.com/bborbe/booking/authentication"
	booking_model "github.com/bborbe/booking/model"
	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
	"github.com/bborbe/server/idparser"
)

var (
	logger = log.DefaultLogger
)

type Get func(int) (*booking_model.Model, error)

type IsParticipant func(authentication *booking_authentication.Authentication) bool

type HttpRequestToAuthentication func(request *http.Request) (*booking_authentication.Authentication, error)

type GetByToken func(token string) (*booking_model.Model, error)

type handler struct {
	get                         Get
	httpRequestToAuthentication HttpRequestToAuthentication
	isParticipant               IsParticipant
	getByToken                  GetByToken
}

func New(get Get, httpRequestToAuthentication HttpRequestToAuthentication, isParticipant IsParticipant, getByToken GetByToken) *handler {
	h := new(handler)
	h.get = get
	h.httpRequestToAuthentication = httpRequestToAuthentication
	h.isParticipant = isParticipant
	h.getByToken = getByToken
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	logger.Debug("get model")
	authentication, err := h.httpRequestToAuthentication(request)
	if err != nil {
		return err
	}
	var model *booking_model.Model
	if h.isParticipant(authentication) {
		model, err = h.getByToken(authentication.Token)
		if err != nil {
			return err
		}
	} else {
		value, err := idparser.ParseIdFormRequest(request)
		if err != nil {
			return err
		}
		model, err = h.get(value)
		if err != nil {
			return err
		}
	}
	j := json_handler.NewJsonHandler(model)
	j.ServeHTTP(responseWriter, request)
	return nil
}
