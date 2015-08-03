package tokengenerator

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestGenerateToken(t *testing.T) {
	s := New()
	to, err := s.GenerateToken()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(to), Gt(0)); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateTokenChange(t *testing.T) {
	s := New()
	t1, _ := s.GenerateToken()
	t2, _ := s.GenerateToken()
	if err := AssertThat(t1, Not(Is(t2))); err != nil {
		t.Fatal(err)
	}
}
