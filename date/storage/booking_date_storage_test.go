package storage

import (
	"testing"

	"time"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/date"
)

func TestImplementsStorage(t *testing.T) {
	r := New(nil)
	var i *Storage
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateDate(t *testing.T) {
	var err error
	var dates *[]date.Date
	storage := New(database.New("/tmp/booking_test.db", true))
	if err = storage.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &date.Date{
		Start: time.Now(),
		End:   time.Now(),
	}

	dates, err = storage.FindDates()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(dates, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*dates), Is(0)); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(storage.CreateDate(d), NilValue()); err != nil {
		t.Fatal(err)
	}
	dates, err = storage.FindDates()
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
