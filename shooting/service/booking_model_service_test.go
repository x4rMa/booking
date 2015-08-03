package service

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsShootingService(t *testing.T) {
	r := New(nil)
	var i *Service
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
