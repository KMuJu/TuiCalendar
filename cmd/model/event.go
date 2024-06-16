package model

import (
	"time"
)

type Event struct {
	Id          string    `json:"id"`
	Date        int       `json:"date"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}

func (e Event) Equal(other Event) bool {
	return e.Id == other.Id &&
		e.Name == other.Name
}
