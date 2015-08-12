package user

import "time"

type User struct {
	Id       int       `json:"id" sql:"AUTO_INCREMENT"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}
