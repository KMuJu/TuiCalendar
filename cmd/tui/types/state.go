package types

import "github.com/kmuju/TuiCalendar/cmd/model"

type State interface {
	Render() string
	Focus()
	FocusLost()
	HandleWidthChange(int)
	HandleHeightChange(int)
}

type ListState interface {
	State

	Len() int
	MoveStart(int)
	Up()
	Down()
	GetSelectedEvent() model.Event
}
