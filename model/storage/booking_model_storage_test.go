package storage

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/model"
)

func createStorage() *storage {
	return New(database.New("/tmp/booking_test.db", true))
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

func TestCreateModel(t *testing.T) {
	var err error
	var models *[]model.Model
	s := createStorage()
	if err = s.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &model.Model{
		FirstName: "Hello",
		LastName:  "World",
		Email:     "test@example.com",
		Phone:     "0123456789",
		Token:     "ABC",
	}

	models, err = s.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(models, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*models), Is(0)); err != nil {
		t.Fatal(err)
	}
	_, err = s.Create(d)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	models, err = s.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(models, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*models), Is(1)); err != nil {
		t.Fatal(err)
	}
}
