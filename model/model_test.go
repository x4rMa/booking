package model

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsModel(t *testing.T) {
	r := New()
	var i *Model
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
