package book

import (
	"testing"

	booking_handler "github.com/bborbe/booking/handler"

	. "github.com/bborbe/assert"
)

func TestImplementsHttpHandler(t *testing.T) {
	r := New(nil, nil, nil, nil)
	var i *booking_handler.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
