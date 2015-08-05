package service

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/database/sqlite"
	"github.com/bborbe/booking/user"
	"github.com/bborbe/booking/user/storage"
)

func createService() *userService {
	return New(storage.New(sqlite.New("/tmp/booking_test.db", true)))
}

func TestImplementsUserService(t *testing.T) {
	r := createService()
	var i *Service
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyLogin(t *testing.T) {
	var err error
	var valid bool
	r := createService()
	username := "testuser"
	password := "pass123"
	_, err = r.Create(&user.User{Login: username, Password: password})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}

	valid, err = r.VerifyLogin(&user.User{Login: "wrongusername", Password: password})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(false)); err != nil {
		t.Fatal(err)
	}

	valid, err = r.VerifyLogin(&user.User{Login: username, Password: "invalidpass"})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(false)); err != nil {
		t.Fatal(err)
	}

	valid, err = r.VerifyLogin(&user.User{Login: username, Password: password})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(true)); err != nil {
		t.Fatal(err)
	}
}
