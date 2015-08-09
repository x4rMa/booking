package book

import (
	"testing"

	booking_error_handler "github.com/bborbe/booking/error_handler"

	. "github.com/bborbe/assert"
)

func TestImplementsHttpHandler(t *testing.T) {
	r := New(nil)
	var i *booking_error_handler.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
