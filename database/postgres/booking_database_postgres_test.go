package postgres

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/database"
)

func TestImplementsDatabase(t *testing.T) {
	r := New("name", "user", "pass", false)
	var i *database.Database
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
