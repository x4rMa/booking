package storage

import (
	"time"

	booking_database "github.com/bborbe/booking/database"
	booking_model "github.com/bborbe/booking/model"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	Find() (*[]booking_model.Model, error)
	Create(model *booking_model.Model) (*booking_model.Model, error)
	Get(id int) (*booking_model.Model, error)
	Delete(id int) (*booking_model.Model, error)
	Update(model *booking_model.Model) (*booking_model.Model, error)
	GetByToken(token string) (*booking_model.Model, error)
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
	err = db.DropTableIfExists(&booking_model.Model{}).Error
	if err != nil {
		return err
	}
	err = db.CreateTable(&booking_model.Model{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) Find() (*[]booking_model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	models := &[]booking_model.Model{}
	query := db.Find(models)
	return models, query.Error
}

func (s *storage) Create(model *booking_model.Model) (*booking_model.Model, error) {
	model.Created = time.Now()
	model.Updated = time.Now()
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Create(model)
	return model, query.Error
}

func (s *storage) Update(model *booking_model.Model) (*booking_model.Model, error) {
	model.Updated = time.Now()
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Save(model)
	return model, query.Error
}

func (s *storage) Get(id int) (*booking_model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	model := &booking_model.Model{}
	query := db.First(model, id)
	if query.Error != nil {
		return nil, err
	}
	return model, nil
}

func (s *storage) Delete(id int) (*booking_model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	model := &booking_model.Model{}
	query := db.First(model, id)
	if query.Error != nil {
		return nil, err
	}
	query = db.Delete(model)
	if query.Error != nil {
		return nil, err
	}
	return model, nil
}

func (s *storage) GetByToken(token string) (*booking_model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	model := &booking_model.Model{}
	query := db.Where(booking_model.Model{Token: token}).First(model)
	if query.Error != nil {
		return nil, err
	}
	return model, nil
}
