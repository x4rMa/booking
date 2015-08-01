package storage

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsStorage(t *testing.T) {
	r := New(nil)
	var i *Storage
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
