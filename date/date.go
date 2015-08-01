package date

import "time"

type Date interface {
	GetStart() time.Time
	SetStart(start time.Time)
	GetEnd() time.Time
	SetEnd(end time.Time)
}

type date struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func New() *date {
	return new(date)
}

func (d *date) SetStart(start time.Time) {
	d.Start = start
}

func (d *date) GetStart() time.Time {
	return d.Start
}

func (d *date) SetEnd(end time.Time) {
	d.End = end
}

func (d *date) GetEnd() time.Time {
	return d.End
}
