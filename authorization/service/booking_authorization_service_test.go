package service

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/authentication"
)

func TestImplementsAuthenticationService(t *testing.T) {
	r := New()
	var i *Service
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestHasRole(t *testing.T) {
	r := New()
	hasRole, err := r.HasRole(&authentication.Authentication{}, "role")
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(true)); err != nil {
		t.Fatal(err)
	}
}
