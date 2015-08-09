package storage

import (
	booking_database "github.com/bborbe/booking/database"
	booking_user "github.com/bborbe/booking/user"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	Find() (*[]booking_user.User, error)
	Create(user *booking_user.User) (*booking_user.User, error)
	Get(id int) (*booking_user.User, error)
	Delete(id int) (*booking_user.User, error)
	Update(user *booking_user.User) (*booking_user.User, error)
	FindByLogin(login string) (*[]booking_user.User, error)
}

type storage struct {
	database booking_database.Database
}

func New(database booking_database.Database) *storage {
	s := new(storage)
	s.database = database
	s.createDefaultUser()
	return s
}

func (s *storage) createDefaultUser() error {
	users, err := s.Find()
	if err != nil {
		return err
	}
	if len(*users) == 0 {
		_, err = s.Create(&booking_user.User{Login: "admin", Password: "test123"})
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
	err = db.DropTableIfExists(&booking_user.User{}).Error
	if err != nil {
		return err
	}
	err = db.CreateTable(&booking_user.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) Find() (*[]booking_user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	users := &[]booking_user.User{}
	query := db.Find(users)
	return users, query.Error
}

func (s *storage) Create(user *booking_user.User) (*booking_user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Create(user)
	return user, query.Error
}

func (s *storage) Update(user *booking_user.User) (*booking_user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	query := db.Save(user)
	return user, query.Error
}

func (s *storage) Get(id int) (*booking_user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	user := &booking_user.User{}
	query := db.First(user, id)
	if query.Error != nil {
		return nil, err
	}
	return user, nil
}

func (s *storage) Delete(id int) (*booking_user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	user := &booking_user.User{}
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

func (s *storage) FindByLogin(login string) (*[]booking_user.User, error) {
	db, err := s.database.DB()
	if err != nil {
		return nil, err
	}
	users := &[]booking_user.User{}
	query := db.Where(booking_user.User{Login: login}).Find(users)
	if query.Error != nil {
		return nil, err
	}
	return users, nil
}
