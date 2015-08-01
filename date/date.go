package date

import "time"

type Date struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
