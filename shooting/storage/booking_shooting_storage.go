package storage

import (
	"time"

	booking_database "github.com/bborbe/booking/database"
	booking_shooting "github.com/bborbe/booking/shooting"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	Find() (*[]booking_shooting.Shooting, error)
	Create(shooting *booking_shooting.Shooting) (*booking_shooting.Shooting, error)
	Get(id int) (*booking_shooting.Shooting, error)
	Delete(id int) (*booking_shooting.Shooting, error)
	Update(shooting *booking_shooting.Shooting) (*booking_shooting.Shooting, error)
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
	err = db.DropTableIfExists(&booking_shooting.Shooting{}).Error
	if err != nil {
		return err
	}
	err = db.CreateTable(&booking_shooting.Shooting{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) Find() (*[]booking_shooting.Shooting, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	shootings := &[]booking_shooting.Shooting{}
	query := db.Find(shootings)
	return shootings, query.Error
}

func (s *storage) Create(shooting *booking_shooting.Shooting) (*booking_shooting.Shooting, error) {
	shooting.Created = time.Now()
	shooting.Updated = time.Now()
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Create(shooting)
	return shooting, query.Error
}

func (s *storage) Update(shooting *booking_shooting.Shooting) (*booking_shooting.Shooting, error) {
	shooting.Updated = time.Now()
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Save(shooting)
	return shooting, query.Error
}

func (s *storage) Get(id int) (*booking_shooting.Shooting, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	shooting := &booking_shooting.Shooting{}
	query := db.First(shooting, id)
	if query.Error != nil {
		return nil, err
	}
	return shooting, nil
}

func (s *storage) Delete(id int) (*booking_shooting.Shooting, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	shooting := &booking_shooting.Shooting{}
	query := db.First(shooting, id)
	if query.Error != nil {
		return nil, err
	}
	query = db.Delete(shooting)
	if query.Error != nil {
		return nil, err
	}
	return shooting, nil
}
