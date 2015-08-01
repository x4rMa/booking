package shooting

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsShooting(t *testing.T) {
	r := New()
	var i *Shooting
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
