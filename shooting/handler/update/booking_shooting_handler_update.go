package update

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	booking_shooting "github.com/bborbe/booking/shooting"

	"github.com/bborbe/log"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type Update func(*booking_shooting.Shooting) (*booking_shooting.Shooting, error)

type handler struct {
	update Update
}

func New(update Update) *handler {
	h := new(handler)
	h.update = update
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	logger.Debugf("shooting update: %s", string(content))
	var f booking_shooting.Shooting
	err = json.Unmarshal(content, &f)
	if err != nil {
		return err
	}
	obj, err := h.update(&f)
	if err != nil {
		return err
	}
	j := json_handler.NewJsonHandler(obj)
	j.ServeHTTP(responseWriter, request)
	return nil
}
