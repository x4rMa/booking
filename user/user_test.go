package user

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsUser(t *testing.T) {
	r := NewUser()
	var i *User
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
