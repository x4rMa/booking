package model

type Model interface {
	GetFirstName() string
	SetFirstName(firstname string)
	GetLastName() string
	SetLastName(lastname string)
	GetEmail() string
	SetEmail(email string)
	GetPhone() string
	SetPhone(phone string)
}

type model struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func New() *model {
	return new(model)
}

func (d *model) GetFirstName() string {
	return d.FirstName
}

func (d *model) SetFirstName(firstname string) {
	d.FirstName = firstname
}

func (d *model) GetLastName() string {
	return d.LastName
}

func (d *model) SetLastName(lastname string) {
	d.LastName = lastname
}

func (d *model) GetEmail() string {
	return d.Email
}

func (d *model) SetEmail(email string) {
	d.Email = email
}

func (d *model) GetPhone() string {
	return d.Phone
}

func (d *model) SetPhone(phone string) {
	d.Phone = phone
}
