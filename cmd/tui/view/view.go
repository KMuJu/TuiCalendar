package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kmuju/TuiCalendar/cmd/model"
	"github.com/kmuju/TuiCalendar/cmd/tui/state"
	"github.com/kmuju/TuiCalendar/cmd/tui/types"
)

type BaseView struct {
	height       int
	selected     int
	eventlist    types.ListState
	eventpreview *state.EventPreview
	focusables   []types.FocusAble
	focusLen     int
}

func NewBaseView(events []model.Event, width, height int) BaseView {
	listwidth := width / 4
	eventlist := state.NewEventList(events, listwidth, height, 0, true)
	eventpreview := state.NewPreviewer(width-listwidth, height)
	focusables := []types.FocusAble{
		eventlist,
		eventpreview,
	}
	base := BaseView{
		height:       height,
		selected:     0,
		eventlist:    eventlist,
		eventpreview: eventpreview,
		focusables:   focusables,
		focusLen:     len(focusables),
	}

	base.updateFocus()

	return base
}

func (self *BaseView) Render() string {
	list := self.eventlist.Render()
	preview := self.eventpreview.Render(self.eventlist.GetSelectedEvent())
	middle := strings.Repeat(" ┃\n", self.height-1)
	return lipgloss.JoinHorizontal(lipgloss.Top, list, middle, preview)
}

func (self *BaseView) updateFocus() {
	self.focusables[self.selected].Focus()
	for i := 0; i < self.focusLen; i++ {
		if self.selected == i {
			continue
		}
		self.focusables[i].FocusLost()
	}
}

func (self *BaseView) HandleKey(key string) {
	switch key {
	case "ctrl+l":
		self.selected++
		if self.selected == self.focusLen {
			self.selected = 0
		}
		self.updateFocus()
		break
	case "ctrl+h":
		if self.selected == 0 {
			self.selected = self.focusLen
		}
		self.selected--
		self.updateFocus()
		break
	default:
		self.eventlist.HandleKey(key)
	}
}
