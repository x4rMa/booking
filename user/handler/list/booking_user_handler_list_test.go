package list

import (
	"testing"
	booking_user "github.com/bborbe/booking/user"


	"net/http"

	. "github.com/bborbe/assert"
)

func TestImplementsHttpHandler(t *testing.T) {
	r := New(func() (*[]booking_user.User, error) {
		return nil, nil
	})
	var i *http.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
