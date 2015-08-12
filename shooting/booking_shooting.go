package shooting

import "time"

type Shooting struct {
	Id      int       `json:"id" sql:"AUTO_INCREMENT"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Name    string    `json:"name"`
	ModelId int       `json:"model_id"`
	DateId  int       `json:"date_id"`
}
