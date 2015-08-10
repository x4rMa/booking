package get

import (
	booking_handler "github.com/bborbe/booking/handler"

	"testing"

	. "github.com/bborbe/assert"
)

func TestNewHandlerImplementsHttpHandler(t *testing.T) {
	r := New(nil)
	var i (*booking_handler.Handler) = nil
	err := AssertThat(r, Implements(i).Message("check implements http.Handler"))
	if err != nil {
		t.Fatal(err)
	}
}
