package user

type userService struct {
}

type UserService interface {
	List() []User
}

func NewUserService() *userService {
	return new(userService)
}

func (u *userService) List() []User {
	users := make([]User, 0)
	users = append(users, createUser("John", "Doe"))
	users = append(users, createUser("Anna", "Smith"))
	users = append(users, createUser("Peter", "Jones"))
	return users
}

func createUser(firstName string, lastName string) User {
	u := NewUser()
	u.FirstName = firstName
	u.LastName = lastName
	return u
}
