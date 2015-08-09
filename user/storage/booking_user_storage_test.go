package storage

import (
	"testing"

	. "github.com/bborbe/assert"
	booking_database_sqlite "github.com/bborbe/booking/database/sqlite"
	booking_user "github.com/bborbe/booking/user"
)

func createStorage() *storage {
	return New(booking_database_sqlite.New("/tmp/booking_test.db", false))
}

func TestImplementsStorage(t *testing.T) {
	s := createStorage()
	var i *Storage
	err := AssertThat(s, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestListEmpty(t *testing.T) {
	s := createStorage()
	list, err := s.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(list, NotNilValue()); err != nil {
		t.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {
	var err error
	var users *[]booking_user.User
	s := createStorage()
	if err = s.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &booking_user.User{
		Login:    "admin",
		Password: "test123",
	}

	users, err = s.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(users, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*users), Is(0)); err != nil {
		t.Fatal(err)
	}
	_, err = s.Create(d)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	users, err = s.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(users, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*users), Is(1)); err != nil {
		t.Fatal(err)
	}
}

func TestCreateUserHasId(t *testing.T) {
	var err error
	s := createStorage()
	if err = s.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &booking_user.User{
		Login:    "admin",
		Password: "test123",
	}
	m, err := s.Create(d)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(m, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(m.Id, Gt(0)); err != nil {
		t.Fatal(err)
	}
}
