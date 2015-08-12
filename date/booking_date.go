package date

import "time"

type Date struct {
	Id      int       `json:"id" sql:"AUTO_INCREMENT"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
}
