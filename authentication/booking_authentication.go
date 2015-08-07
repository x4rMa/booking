package authentication

type Authentication struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
