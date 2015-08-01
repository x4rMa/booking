package database

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsDatabase(t *testing.T) {
	r := New("/tmp", false)
	var i *Database
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
