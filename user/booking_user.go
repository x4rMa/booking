package user


type User struct {
	Id       int    `json:"id" sql:"AUTO_INCREMENT"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
