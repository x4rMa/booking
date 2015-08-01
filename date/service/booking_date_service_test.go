package service

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsDateService(t *testing.T) {
	r := New(nil)
	var i *DateService
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
