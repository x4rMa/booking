package model

type Model struct {
	Id        int    `json:"id" sql:"AUTO_INCREMENT"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Token     string `json:"token" sql:"type:varchar(255);unique_index`
}
