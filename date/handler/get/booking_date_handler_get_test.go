package get

import (
	"testing"

	booking_error_handler "github.com/bborbe/booking/error_handler"

	. "github.com/bborbe/assert"
)

func TestNewHandlerImplementsHttpHandler(t *testing.T) {
	r := New(nil)
	var i (*booking_error_handler.Handler) = nil
	err := AssertThat(r, Implements(i).Message("check implements http.Handler"))
	if err != nil {
		t.Fatal(err)
	}
}
