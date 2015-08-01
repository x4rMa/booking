package service

import "github.com/bborbe/booking/user"

type userService struct {
}

type UserService interface {
	List() []*user.User
}

func NewUserService() *userService {
	return new(userService)
}

func (u *userService) List() []*user.User {
	users := make([]*user.User, 0)
	users = append(users, createUser("John", "Doe"))
	users = append(users, createUser("Anna", "Smith"))
	users = append(users, createUser("Peter", "Jones"))
	return users
}

func createUser(firstName string, lastName string) *user.User {
	u := &user.User{}
	u.FirstName = firstName
	u.LastName = lastName
	return u
}
