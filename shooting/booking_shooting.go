package shooting

type Shooting struct {
	Id      int    `json:"id" sql:"AUTO_INCREMENT"`
	Name    string `json:"name"`
	ModelId int    `json:"model_id"`
	DateId  int    `json:"date_id"`
}
