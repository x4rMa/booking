package booked_event

import "github.com/bborbe/booking/shooting"

type BookedEvent struct {
	Shooting shooting.Shooting
}

func New(shooting shooting.Shooting) *BookedEvent {
	e := new(BookedEvent)
	e.Shooting = shooting
	return e
}
