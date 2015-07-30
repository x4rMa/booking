package service

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsUserService(t *testing.T) {
	r := NewUserService()
	var i *UserService
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserServiceList(t *testing.T) {
	r := NewUserService()
	users := r.List()
	err := AssertThat(users, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(users), Is(3))
	if err != nil {
		t.Fatal(err)
	}
}
