package storage

import (
	"fmt"

	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/model"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	Find() (*[]model.Model, error)
	Create(model *model.Model) (*model.Model, error)
	Get(id int) (*model.Model, error)
	Delete(id int) (*model.Model, error)
	Update(model *model.Model) (*model.Model, error)
	FindByToken(token string) (*[]model.Model, error)
}

type storage struct {
	database database.Database
}

func New(database database.Database) *storage {
	s := new(storage)
	s.database = database
	db, err := s.database.DB()
	if err != nil {
		panic(fmt.Sprintf("auto migrate failed: %v", err))
	}
	db.AutoMigrate(&model.Model{})
	return s
}

func (s *storage) Truncate() error {
	db, err := s.database.DB()
	if err != nil {
		return err
	}
	err = db.DropTableIfExists(&model.Model{}).Error
	if err != nil {
		return err
	}
	err = db.CreateTable(&model.Model{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) Find() (*[]model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	models := &[]model.Model{}
	query := db.Find(models)
	return models, query.Error
}

func (s *storage) Create(model *model.Model) (*model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Create(model)
	return model, query.Error
}

func (s *storage) Update(model *model.Model) (*model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Save(model)
	return model, query.Error
}

func (s *storage) Get(id int) (*model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	model := &model.Model{}
	query := db.First(model, id)
	if query.Error != nil {
		return nil, err
	}
	return model, nil
}

func (s *storage) Delete(id int) (*model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	model := &model.Model{}
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

func (s *storage) FindByToken(token string) (*[]model.Model, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	models := &[]model.Model{}
	query := db.Where(model.Model{Token: token}).Find(models)
	if query.Error != nil {
		return nil, err
	}
	return models, nil
}
