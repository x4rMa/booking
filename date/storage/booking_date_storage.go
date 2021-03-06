package storage

import (
	"time"

	booking_database "github.com/bborbe/booking/database"
	booking_date "github.com/bborbe/booking/date"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	Find() (*[]booking_date.Date, error)
	FindWithoutShootingAndInFuture() (*[]booking_date.Date, error)
	Create(date *booking_date.Date) (*booking_date.Date, error)
	Get(id int) (*booking_date.Date, error)
	Delete(id int) (*booking_date.Date, error)
	Update(date *booking_date.Date) (*booking_date.Date, error)
}

type storage struct {
	database booking_database.Database
}

func New(database booking_database.Database) *storage {
	s := new(storage)
	s.database = database
	return s
}

func (s *storage) Truncate() error {
	db, err := s.database.DB()
	if err != nil {
		return err
	}
	err = db.DropTableIfExists(&booking_date.Date{}).Error
	if err != nil {
		return err
	}
	err = db.CreateTable(&booking_date.Date{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) Find() (*[]booking_date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	dates := &[]booking_date.Date{}
	query := db.Find(dates)
	return dates, query.Error
}

func (s *storage) FindWithoutShootingAndInFuture() (*[]booking_date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	dates := &[]booking_date.Date{}
	query := db.Find(dates)
	db.Joins("LEFT JOIN shooting ON shooting.date_id = date.id").Where("shooting.id IS NULL AND DATE(date.start) > DATE(?)", time.Now()).Find(dates)
	return dates, query.Error
}

func (s *storage) Create(date *booking_date.Date) (*booking_date.Date, error) {
	date.Created = time.Now()
	date.Updated = time.Now()
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Create(date)
	return date, query.Error
}

func (s *storage) Update(date *booking_date.Date) (*booking_date.Date, error) {
	date.Updated = time.Now()
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Save(date)
	return date, query.Error
}

func (s *storage) Get(id int) (*booking_date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	date := &booking_date.Date{}
	query := db.First(date, id)
	if query.Error != nil {
		return nil, err
	}
	return date, nil
}

func (s *storage) Delete(id int) (*booking_date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	date := &booking_date.Date{}
	query := db.First(date, id)
	if query.Error != nil {
		return nil, err
	}
	query = db.Delete(date)
	if query.Error != nil {
		return nil, err
	}
	return date, nil
}
