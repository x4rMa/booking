package storage

import (
	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/date"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	FindDates() (*[]date.Date, error)
	FindLatestDates(limit int) (*[]date.Date, error)
	CreateDate(date *date.Date) error
	GetDate(id int) (*date.Date, error)
	DeleteDate(id int) (*date.Date, error)
}

type storage struct {
	database database.Database
}

func New(database database.Database) *storage {
	s := new(storage)
	s.database = database
	return s
}

func (s *storage) Truncate() error {
	db, err := s.database.DB()
	if err != nil {
		return err
	}
	err = db.DropTableIfExists(&date.Date{}).Error
	if err != nil {
		return err
	}
	err = db.CreateTable(&date.Date{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) FindDates() (*[]date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	dates := &[]date.Date{}
	query := db.Order("timestamp desc").Find(dates)
	return dates, query.Error
}

func (s *storage) FindLatestDates(limit int) (*[]date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	dates := &[]date.Date{}
	query := db.Order("timestamp desc").Limit(limit).Find(dates)
	return dates, query.Error
}

func (s *storage) CreateDate(date *date.Date) error {
	db, err := s.database.DB()
	if err != nil {
		return err
	}
	query := db.Create(date)
	return query.Error
}

func (s *storage) GetDate(id int) (*date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	date := &date.Date{}
	query := db.First(date, id)
	if query.Error != nil {
		return nil, err
	}
	return date, nil
}

func (s *storage) DeleteDate(id int) (*date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	date := &date.Date{}
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
