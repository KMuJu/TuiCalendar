package tui

import "github.com/kmuju/TuiCalendar/cmd/model"

type Calendar struct {
	events       []model.Event
	height       int
	width        int
	listWidth    int
	renderFrom   int
	renderAmount int
	selected     int
}
