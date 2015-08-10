package mock

import (
	"testing"

	. "github.com/bborbe/assert"
	booking_handler "github.com/bborbe/booking/handler"
)

func TestImplementsHttpHandler(t *testing.T) {
	r := New()
	var i *booking_handler.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
