package status_handler

import (
	"net/http"

	"fmt"
)

type handler struct {
}

func New() *handler {
	h := new(handler)
	return h
}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) error {
	fmt.Fprint(resp, "OK")
	return nil
}
