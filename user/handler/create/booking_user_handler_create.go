package create

import (
	booking_user "github.com/bborbe/booking/user"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bborbe/log"
	error_handler "github.com/bborbe/server/handler/error"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type Create func( *booking_user.User) (*booking_user.User, error)

type handler struct {
	create Create
}

func New(create Create) *handler {
	h := new(handler)
	h.create = create
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	err := h.serveHTTP(responseWriter, request)
	if err != nil {
		logger.Debug(err)
		e := error_handler.NewErrorMessage(http.StatusInternalServerError, err.Error())
		e.ServeHTTP(responseWriter, request)
		return
	}
}

func (h *handler) serveHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	logger.Debugf("user create: %s", string(content))
	var f booking_user.User
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
