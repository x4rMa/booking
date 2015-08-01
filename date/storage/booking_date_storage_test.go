package storage

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsStorage(t *testing.T) {
	r := New("/tmp/stats_test.db", false)
	var i *Storage
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
