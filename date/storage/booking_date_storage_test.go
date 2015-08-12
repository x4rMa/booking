package storage

import (
	"testing"

	"time"

	. "github.com/bborbe/assert"
	booking_database "github.com/bborbe/booking/database"
	booking_database_sqlite "github.com/bborbe/booking/database/sqlite"
	booking_date "github.com/bborbe/booking/date"
	booking_shooting "github.com/bborbe/booking/shooting"
	shooting_storage "github.com/bborbe/booking/shooting/storage"
)

func createDatabase() booking_database.Database {
	return booking_database_sqlite.New("/tmp/booking_test.db", true)
}

func createStorage(db booking_database.Database) *storage {
	return New(db)
}

func TestImplementsStorage(t *testing.T) {
	dateStorage := createStorage(createDatabase())
	var i *Storage
	err := AssertThat(dateStorage, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestListEmpty(t *testing.T) {
	dateStorage := createStorage(createDatabase())
	list, err := dateStorage.Find()
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
	dateStorage := createStorage(createDatabase())
	if err = dateStorage.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &booking_date.Date{
		Start: time.Now(),
		End:   time.Now(),
	}

	dates, err = dateStorage.Find()
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(dates, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(*dates), Is(0)); err != nil {
		t.Fatal(err)
	}
	_, err = dateStorage.Create(d)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	dates, err = dateStorage.Find()
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
	dateStorage := createStorage(createDatabase())
	if err = dateStorage.Truncate(); err != nil {
		t.Fatal(err)
	}
	d := &booking_date.Date{
		Start: time.Now(),
		End:   time.Now(),
	}
	m, err := dateStorage.Create(d)
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

func TestFindWithoutShootingAndInFuture(t *testing.T) {
	var err error
	var dates *[]booking_date.Date
	database := createDatabase()
	dateStorage := createStorage(database)
	shootingStorage := shooting_storage.New(database)
	if err = dateStorage.Truncate(); err != nil {
		t.Fatal(err)
	}
	// no dates
	{
		dates, err = dateStorage.FindWithoutShootingAndInFuture()
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
			Start: time.Now().Add(24 * time.Hour),
			End:   time.Now().Add(25 * time.Hour),
		}
		_, err = dateStorage.Create(d)
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		dates, err = dateStorage.FindWithoutShootingAndInFuture()
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
			Start: time.Now().Add(24 * time.Hour),
			End:   time.Now().Add(25 * time.Hour),
		}
		d, err = dateStorage.Create(d)
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		_, err = shootingStorage.Create(&booking_shooting.Shooting{Name: "test", DateId: d.Id})
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		dates, err = dateStorage.FindWithoutShootingAndInFuture()
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		if err = AssertThat(len(*dates), Is(1)); err != nil {
			t.Fatal(err)
		}
	}
	// three dates one with shooting
	{

		d := &booking_date.Date{
			Start: time.Now().Add(-2 * time.Hour),
			End:   time.Now().Add(-1 * time.Hour),
		}
		d, err = dateStorage.Create(d)
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		dates, err = dateStorage.FindWithoutShootingAndInFuture()
		if err = AssertThat(err, NilValue()); err != nil {
			t.Fatal(err)
		}
		if err = AssertThat(len(*dates), Is(1)); err != nil {
			t.Fatal(err)
		}
	}
}
