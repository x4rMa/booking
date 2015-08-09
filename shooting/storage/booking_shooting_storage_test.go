package storage

import (
	"testing"

	. "github.com/bborbe/assert"
	booking_database_sqlite "github.com/bborbe/booking/database/sqlite"
	booking_shooting "github.com/bborbe/booking/shooting"
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

func TestCreateShooting(t *testing.T) {
	var err error
	var shootings *[]booking_shooting.Shooting
	s := createStorage()
	if err = s.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &booking_shooting.Shooting{
		Name: "Hello World",
	}

	shootings, err = s.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(shootings, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*shootings), Is(0)); err != nil {
		t.Fatal(err)
	}
	_, err = s.Create(d)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	shootings, err = s.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(shootings, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*shootings), Is(1)); err != nil {
		t.Fatal(err)
	}
}

func TestCreateShootingHasId(t *testing.T) {
	var err error
	s := createStorage()
	if err = s.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &booking_shooting.Shooting{
		Name: "Hello World",
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
