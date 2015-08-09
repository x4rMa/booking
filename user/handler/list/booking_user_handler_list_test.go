package list

import (
	"testing"

	booking_user "github.com/bborbe/booking/user"

	booking_error_handler "github.com/bborbe/booking/error_handler"

	. "github.com/bborbe/assert"
)

func TestImplementsHttpHandler(t *testing.T) {
	r := New(func() (*[]booking_user.User, error) {
		return nil, nil
	})
	var i *booking_error_handler.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
