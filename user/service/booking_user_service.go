package service

import (
	booking_user "github.com/bborbe/booking/user"
	booking_user_storage "github.com/bborbe/booking/user/storage"
	"github.com/bborbe/log"
	_ "github.com/lib/pq"
)

var (
	logger = log.DefaultLogger
)

type Service interface {
	List() (*[]booking_user.User, error)
	Get(id int) (*booking_user.User, error)
	Create(user *booking_user.User) (*booking_user.User, error)
	Delete(id int) (*booking_user.User, error)
	Update(d *booking_user.User) (*booking_user.User, error)
	VerifyLogin(d *booking_user.User) (bool, error)
}

type userService struct {
	storage booking_user_storage.Storage
}

func New(storage booking_user_storage.Storage) *userService {
	d := new(userService)
	d.storage = storage
	return d
}

func (s *userService) Create(d *booking_user.User) (*booking_user.User, error) {
	logger.Debug("create")
	return s.storage.Create(d)
}

func (s *userService) Update(d *booking_user.User) (*booking_user.User, error) {
	logger.Debug("update")
	return s.storage.Update(d)
}

func (s *userService) List() (*[]booking_user.User, error) {
	logger.Debug("list")
	return s.storage.Find()
}

func (s *userService) Get(id int) (*booking_user.User, error) {
	logger.Debug("get")
	return s.storage.Get(id)
}

func (s *userService) Delete(id int) (*booking_user.User, error) {
	logger.Debug("delete")
	return s.storage.Delete(id)
}

func (s *userService) VerifyLogin(d *booking_user.User) (bool, error) {
	logger.Debugf("verify login for user: %s", d.Login)
	users, err := s.storage.FindByLogin(d.Login)
	if err != nil {
		return false, err
	}
	for _, u := range *users {
		if u.Login == d.Login && u.Password == d.Password {
			return true, nil
		}
	}
	return false, nil
}
