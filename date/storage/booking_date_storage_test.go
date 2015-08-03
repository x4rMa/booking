package storage

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/date"
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

func TestCreateDate(t *testing.T) {
	var err error
	var dates *[]date.Date
	s := createStorage()
	if err = s.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &date.Date{
		Start: "1",
		End:   "2",
	}

	dates, err = s.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(dates, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*dates), Is(0)); err != nil {
		t.Fatal(err)
	}
	_, err = s.Create(d)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	dates, err = s.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(dates, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*dates), Is(1)); err != nil {
		t.Fatal(err)
	}
}
