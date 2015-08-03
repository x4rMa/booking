package get

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
)

func TestNewHandlerImplementsHttpHandler(t *testing.T) {
	r := New(nil)
	var i (*http.Handler) = nil
	err := AssertThat(r, Implements(i).Message("check implements http.Handler"))
	if err != nil {
		t.Fatal(err)
	}
}
