package storage

import (
	"testing"

	. "github.com/bborbe/assert"
	booking_database "github.com/bborbe/booking/database"
	booking_database_sqlite "github.com/bborbe/booking/database/sqlite"
	booking_date "github.com/bborbe/booking/date"
	booking_shooting "github.com/bborbe/booking/shooting"
	shooting_storage "github.com/bborbe/booking/shooting/storage"
)

func createDatabase() booking_database.Database {
	return booking_database_sqlite.New("/tmp/booking_test.db", false)
}

func createStorage(db booking_database.Database) *storage {
	return New(db)
}

func TestImplementsStorage(t *testing.T) {
	s := createStorage(createDatabase())
	var i *Storage
	err := AssertThat(s, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestListEmpty(t *testing.T) {
	s := createStorage(createDatabase())
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
	var dates *[]booking_date.Date
	s := createStorage(createDatabase())
	if err = s.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &booking_date.Date{
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

func TestCreateDateHasId(t *testing.T) {
	var err error
	s := createStorage(createDatabase())
	if err = s.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &booking_date.Date{
		Start: "1",
		End:   "2",
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

func TestFindWithoutShooting(t *testing.T) {
	var err error
	var dates *[]booking_date.Date
	db := createDatabase()
	s := createStorage(db)
	shootingStorage := shooting_storage.New(db)
	if err = s.Truncate(); err != nil {
		t.Fatal(err)
	}
	// no dates
	{
		dates, err = s.FindWithoutShooting()
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		if err = AssertThat(len(*dates), Is(0)); err != nil {
			t.Fatal(err)
		}
	}
	// one date without shooting
	{
		d := &booking_date.Date{
			Start: "1",
			End:   "2",
		}
		_, err = s.Create(d)
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		dates, err = s.FindWithoutShooting()
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		if err = AssertThat(len(*dates), Is(1)); err != nil {
			t.Fatal(err)
		}
	}
	// two dates one with shooting
	{

		d := &booking_date.Date{
			Start: "1",
			End:   "2",
		}
		d, err = s.Create(d)
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		_, err = shootingStorage.Create(&booking_shooting.Shooting{Name: "test", DateId: d.Id})
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		dates, err = s.FindWithoutShooting()
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		if err = AssertThat(len(*dates), Is(1)); err != nil {
			t.Fatal(err)
		}
	}
}
