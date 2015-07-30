package handler

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
)

func TestNewHandlerImplementsHttpHandler(t *testing.T) {
	r := NewHandler("/tmp")
	var i (*http.Handler) = nil
	err := AssertThat(r, Implements(i).Message("check implements server.Server"))
	if err != nil {
		t.Fatal(err)
	}
}
