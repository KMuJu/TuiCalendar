package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kmuju/TuiCalendar/cmd/model"
	"github.com/kmuju/TuiCalendar/cmd/tui/state"
)

type BaseView struct {
	height       int
	selected     int
	eventlist    state.EventListState
	eventpreview state.EventPreview
}

func NewBaseView(events []model.Event, width, height int) BaseView {
	listwidth := width / 4
	return BaseView{
		height:       height,
		selected:     0,
		eventlist:    state.NewEventList(events, listwidth, height, 0, true),
		eventpreview: state.NewPreviewer(width-listwidth, height),
	}
}

func (self *BaseView) Render() string {
	list := self.eventlist.Render()
	preview := self.eventpreview.Render(self.eventlist.GetSelectedEvent())
	middle := strings.Repeat(" ┃\n", self.height)
	return lipgloss.JoinHorizontal(lipgloss.Top, list, middle, preview)
}

func (self *BaseView) HandleKey(key string) {
	switch key {
	default:
		self.eventlist.HandleKey(key)
	}
}
