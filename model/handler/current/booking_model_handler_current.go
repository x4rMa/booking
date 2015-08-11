package current

import (
	"net/http"

	booking_authentication                         "github.com/bborbe/booking/authentication"
	booking_model                        "github.com/bborbe/booking/model"
	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type HttpRequestToAuthentication func(request *http.Request) (*booking_authentication.Authentication, error)
type GetByToken func(token string) (*booking_model.Model, error)

type handler struct {
	httpRequestToAuthentication HttpRequestToAuthentication
	getByToken                  GetByToken
}

func New(httpRequestToAuthentication HttpRequestToAuthentication, getByToken GetByToken) *handler {
	h := new(handler)
	h.httpRequestToAuthentication = httpRequestToAuthentication
	h.getByToken = getByToken
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	logger.Debug("get current model")
	authentication, err := h.httpRequestToAuthentication(request)
	if err != nil {
		return err
	}
	obj, err := h.getByToken(authentication.Token)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
