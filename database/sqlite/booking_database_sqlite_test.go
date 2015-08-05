package sqlite

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/database"
)

func TestImplementsDatabase(t *testing.T) {
	r := New("/tmp", false)
	var i *database.Database
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
