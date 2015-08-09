package service

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsAuthenticationService(t *testing.T) {
	r := New()
	var i *Service
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestHasPermission(t *testing.T) {
	r := New()
	hasPermission, err := r.HasPermission("asdf")
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasPermission, Is(true)); err != nil {
		t.Fatal(err)
	}
}
