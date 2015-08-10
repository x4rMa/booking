package list

import (
	"testing"

	booking_user "github.com/bborbe/booking/user"

	booking_handler "github.com/bborbe/booking/handler"

	. "github.com/bborbe/assert"
)

func createList() List {
	return func() (*[]booking_user.User, error) {
		return nil, nil
	}
}
func TestImplementsHttpHandler(t *testing.T) {
	r := New(createList())
	var i *booking_handler.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
