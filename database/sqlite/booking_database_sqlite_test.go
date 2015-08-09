package sqlite

import (
	"testing"

	. "github.com/bborbe/assert"
	booking_database "github.com/bborbe/booking/database"
)

func TestImplementsDatabase(t *testing.T) {
	r := New("/tmp", false)
	var i *booking_database.Database
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
