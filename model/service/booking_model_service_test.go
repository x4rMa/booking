package service

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsModelService(t *testing.T) {
	r := New(nil, nil)
	var i *Service
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
