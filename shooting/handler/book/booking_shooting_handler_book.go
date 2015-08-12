package book

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fmt"

	booking_authentication "github.com/bborbe/booking/authentication"
	booking_model "github.com/bborbe/booking/model"
	booking_shooting "github.com/bborbe/booking/shooting"
	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type Book func(*booking_shooting.Shooting) (*booking_shooting.Shooting, error)
type HttpRequestToAuthentication func(request *http.Request) (*booking_authentication.Authentication, error)
type GetByToken func(token string) (*booking_model.Model, error)
type Get func(int) (*booking_shooting.Shooting, error)

type handler struct {
	book                        Book
	httpRequestToAuthentication HttpRequestToAuthentication
	getByToken                  GetByToken
	get                         Get
}

func New(httpRequestToAuthentication HttpRequestToAuthentication, getByToken GetByToken, book Book, get Get) *handler {
	h := new(handler)
	h.httpRequestToAuthentication = httpRequestToAuthentication
	h.getByToken = getByToken
	h.book = book
	h.get = get
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
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	logger.Debugf("shooting book: %s", string(content))
	var f booking_shooting.Shooting
	err = json.Unmarshal(content, &f)
	if err != nil {
		return err
	}
	shooting, err := h.get(f.Id)
	if err != nil {
		return err
	}
	if shooting.ModelId != model.Id {
		logger.Debugf("permission %d != %d", f.ModelId, model.Id)
		return fmt.Errorf("permission denied")
	}
	obj, err := h.book(&f)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
