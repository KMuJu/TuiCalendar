package tui

import (
	"time"
)

type Event struct {
	Date        int
	Name        string
	Description string
	Start       time.Time
	End         time.Time
}

type Calendar struct {
	events       []Event
	height       int
	width        int
	listWidth    int
	renderFrom   int
	renderAmount int
	selected     int
}
