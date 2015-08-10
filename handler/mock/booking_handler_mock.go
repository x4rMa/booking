package mock

import "net/http"

type handler struct {
	Response http.ResponseWriter
	Request  *http.Request
	Error    error
}

func New() *handler {
	return new(handler)
}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) error {
	h.Request = req
	h.Response = resp
	return h.Error
}
