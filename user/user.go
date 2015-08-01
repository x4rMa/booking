package user

type User interface {
	GetLastName() string
	SetLastName(name string)
	GetFirstName() string
	SetFirstName(name string)
}

type user struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}

func New() *user {
	return new(user)
}

func (u *user) GetFirstName() string {
	return u.FirstName
}

func (u *user) SetFirstName(name string) {
	u.FirstName = name
}

func (u *user) GetLastName() string {
	return u.LastName
}

func (u *user) SetLastName(name string) {
	u.LastName = name
}
