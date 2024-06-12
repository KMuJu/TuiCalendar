package types

import "github.com/kmuju/TuiCalendar/cmd/model"

type FocusAble interface {
	InnerFocus(int) bool // Wether or not the struct has multiple focuses inside in direction specified
	Focus()
	FocusLost()
	HandleKey(string)
}

type State interface {
	Render() string
	HandleWidthChange(int)
	HandleHeightChange(int)
	HandleKey(string)
}

type ListState interface {
	State
	Len() int
	MoveStart(int)
	Up()
	Down()
	GetSelectedEvent() model.Event
}
