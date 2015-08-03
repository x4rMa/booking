package storage

import (
	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/shooting"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	Find() (*[]shooting.Shooting, error)
	Create(shooting *shooting.Shooting) (*shooting.Shooting, error)
	Get(id int) (*shooting.Shooting, error)
	Delete(id int) (*shooting.Shooting, error)
	Update(shooting *shooting.Shooting) (*shooting.Shooting, error)
}

type storage struct {
	database database.Database
}

func New(database database.Database) *storage {
	s := new(storage)
	s.database = database
	db,err := s.database.DB()
	if err != nil {
		panic("auto migrate failed")
	}
	db.AutoMigrate(&shooting.Shooting{})
	return s
}

func (s *storage) Truncate() error {
	db, err := s.database.DB()
	if err != nil {
		return err
	}
	err = db.DropTableIfExists(&shooting.Shooting{}).Error
	if err != nil {
		return err
	}
	err = db.CreateTable(&shooting.Shooting{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) Find() (*[]shooting.Shooting, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	shootings := &[]shooting.Shooting{}
	query := db.Find(shootings)
	return shootings, query.Error
}

func (s *storage) Create(shooting *shooting.Shooting) (*shooting.Shooting, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Create(shooting)
	return shooting, query.Error
}

func (s *storage) Update(shooting *shooting.Shooting) (*shooting.Shooting, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Save(shooting)
	return shooting, query.Error
}

func (s *storage) Get(id int) (*shooting.Shooting, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	shooting := &shooting.Shooting{}
	query := db.First(shooting, id)
	if query.Error != nil {
		return nil, err
	}
	return shooting, nil
}

func (s *storage) Delete(id int) (*shooting.Shooting, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	shooting := &shooting.Shooting{}
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
