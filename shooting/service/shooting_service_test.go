package service

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsShootingService(t *testing.T) {
	r := New()
	var i *ShootingService
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
