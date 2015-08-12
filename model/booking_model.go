package model

import "time"

type Model struct {
	Id        int       `json:"id" sql:"AUTO_INCREMENT"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Token     string    `json:"token" sql:"type:varchar(255);unique_index`
}
