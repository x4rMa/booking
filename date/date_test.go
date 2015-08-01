package date

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsDate(t *testing.T) {
	r := New()
	var i *Date
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
