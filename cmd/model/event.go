package model

import (
	"time"
)

type Event struct {
	Id          string
	Date        int
	Name        string
	Description string
	Status      string
	Start       time.Time
	End         time.Time
}

func (e Event) Equal(other Event) bool {
	return e.Id == other.Id &&
		e.Name == other.Name
}
