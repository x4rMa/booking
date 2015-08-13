package get

import (
	"testing"

	booking_handler "github.com/bborbe/booking/handler"

	. "github.com/bborbe/assert"
)

func TestNewHandlerImplementsHttpHandler(t *testing.T) {
	r := New(nil, nil, nil, nil)
	var i (*booking_handler.Handler) = nil
	err := AssertThat(r, Implements(i).Message("check implements http.Handler"))
	if err != nil {
		t.Fatal(err)
	}
}
