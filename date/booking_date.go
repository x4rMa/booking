package date

import "time"

type Date struct {
	Id    int       `json:"id"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
