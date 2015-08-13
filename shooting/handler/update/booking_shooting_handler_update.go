package update

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

type Update func(*booking_shooting.Shooting) (*booking_shooting.Shooting, error)
type HttpRequestToAuthentication func(request *http.Request) (*booking_authentication.Authentication, error)
type GetByToken func(token string) (*booking_model.Model, error)
type Get func(int) (*booking_shooting.Shooting, error)
type IsParticipant func(authentication *booking_authentication.Authentication) bool

type handler struct {
	update                      Update
	httpRequestToAuthentication HttpRequestToAuthentication
	getByToken                  GetByToken
	get                         Get
	isParticipant               IsParticipant
}

func New(httpRequestToAuthentication HttpRequestToAuthentication, getByToken GetByToken, update Update, get Get, isParticipant IsParticipant) *handler {
	h := new(handler)
	h.httpRequestToAuthentication = httpRequestToAuthentication
	h.getByToken = getByToken
	h.update = update
	h.get = get
	h.isParticipant = isParticipant
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
	if h.isParticipant(authentication) {
		shooting, err := h.get(f.Id)
		if err != nil {
			return err
		}
		if shooting.ModelId != model.Id {
			logger.Debugf("permission %d != %d", f.ModelId, model.Id)
			return fmt.Errorf("permission denied")
		}
	}
	obj, err := h.update(&f)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
