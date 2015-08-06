package service

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestNewHandlerImplementsMailService(t *testing.T) {
	r := New()
	var i (*MailService) = nil
	err := AssertThat(r, Implements(i).Message("check implements server.Server"))
	if err != nil {
		t.Fatal(err)
	}
}
