package tui

import (
	"time"
)

type Day struct {
	Events       []Event
	NrEvents     int
	Date         int // 1-31
	Month        int // 1-12
	width        int
	height       int
	IsSelected   bool
	Selected     int
	renderfrom   int
	renderamount int
}

type Week struct {
	Days     [7]Day
	Nr       int // 1-52
	width    int
	height   int
	selected int
}

type Event struct {
	Date        int
	Description string
	Name        string
	Start       time.Time
	End         time.Time
}
