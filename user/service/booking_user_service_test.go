package service

import (
	"testing"

	. "github.com/bborbe/assert"
	booking_database_sqlite "github.com/bborbe/booking/database/sqlite"
	booking_user "github.com/bborbe/booking/user"
	booking_user_storage "github.com/bborbe/booking/user/storage"
)

func createService() *userService {
	return New(booking_user_storage.New(booking_database_sqlite.New("/tmp/booking_test.db", false)))
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
	_, err = r.Create(&booking_user.User{Login: username, Password: password})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}

	valid, err = r.VerifyLogin(&booking_user.User{Login: "wrongusername", Password: password})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(false)); err != nil {
		t.Fatal(err)
	}

	valid, err = r.VerifyLogin(&booking_user.User{Login: username, Password: "invalidpass"})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(false)); err != nil {
		t.Fatal(err)
	}

	valid, err = r.VerifyLogin(&booking_user.User{Login: username, Password: password})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(true)); err != nil {
		t.Fatal(err)
	}
}
