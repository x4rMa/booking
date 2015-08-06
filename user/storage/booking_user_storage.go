package storage

import (
	"fmt"

	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/user"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	Find() (*[]user.User, error)
	Create(user *user.User) (*user.User, error)
	Get(id int) (*user.User, error)
	Delete(id int) (*user.User, error)
	Update(user *user.User) (*user.User, error)
	FindByLogin(login string) (*[]user.User, error)
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
	db.AutoMigrate(&user.User{})
	s.createDefaultUser()
	return s
}

func (s *storage) createDefaultUser() error {
	users, err := s.Find()
	if err != nil {
		return err
	}
	if len(*users) == 0 {
		_, err = s.Create(&user.User{Login: "admin", Password: "test123"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *storage) Truncate() error {
	db, err := s.database.DB()
	if err != nil {
		return err
	}
	err = db.DropTableIfExists(&user.User{}).Error
	if err != nil {
		return err
	}
	err = db.CreateTable(&user.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) Find() (*[]user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	users := &[]user.User{}
	query := db.Find(users)
	return users, query.Error
}

func (s *storage) Create(user *user.User) (*user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Create(user)
	return user, query.Error
}

func (s *storage) Update(user *user.User) (*user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Save(user)
	return user, query.Error
}

func (s *storage) Get(id int) (*user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	user := &user.User{}
	query := db.First(user, id)
	if query.Error != nil {
		return nil, err
	}
	return user, nil
}

func (s *storage) Delete(id int) (*user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	user := &user.User{}
	query := db.First(user, id)
	if query.Error != nil {
		return nil, err
	}
	query = db.Delete(user)
	if query.Error != nil {
		return nil, err
	}
	return user, nil
}

func (s *storage) FindByLogin(login string) (*[]user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	users := &[]user.User{}
	query := db.Where(user.User{Login: login}).Find(users)
	if query.Error != nil {
		return nil, err
	}
	return users, nil
}
