package storage

import (
	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/date"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	Find() (*[]date.Date, error)
	FindWithoutShooting() (*[]date.Date, error)
	Create(date *date.Date) (*date.Date, error)
	Get(id int) (*date.Date, error)
	Delete(id int) (*date.Date, error)
	Update(date *date.Date) (*date.Date, error)
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

func (s *storage) Find() (*[]date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	dates := &[]date.Date{}
	query := db.Find(dates)
	return dates, query.Error
}

func (s *storage) FindWithoutShooting() (*[]date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	dates := &[]date.Date{}
	query := db.Find(dates)
	db.Joins("LEFT JOIN shooting ON shooting.date_id = date.id").Where("shooting.id IS NULL").Find(dates)
	return dates, query.Error
}

func (s *storage) Create(date *date.Date) (*date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Create(date)
	return date, query.Error
}

func (s *storage) Update(date *date.Date) (*date.Date, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Save(date)
	return date, query.Error
}

func (s *storage) Get(id int) (*date.Date, error) {
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

func (s *storage) Delete(id int) (*date.Date, error) {
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
