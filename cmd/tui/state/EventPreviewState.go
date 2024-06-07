package state

import (
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	"github.com/kmuju/TuiCalendar/cmd/model"
)

type ListPreview struct {
	event  model.Event
	width  int
	height int
	focus  bool
}

func (self *ListPreview) SetEvent(event model.Event) {
	self.event = event
}

func (self *ListPreview) Focus()     { self.focus = true }
func (self *ListPreview) FocusLost() { self.focus = true }

func (self *ListPreview) HandleWidthChange(delta int) {
	if self.width+delta > 0 {
		self.width += delta
	}
}

func (self *ListPreview) HandleHeightChange(delta int) {
	if self.height+delta > 0 {
		self.height += delta
	}
}

func (self *ListPreview) Render() string {
	namestyle := lipgloss.NewStyle().Inherit(namestyle)
	datestyle := lipgloss.NewStyle().Inherit(datestyle)

	if self.focus {
		namestyle.BorderForeground(selectedColor)
		datestyle.BorderForeground(selectedColor)
	}

	event := self.event
	width := self.width
	name := namestyle.
		Width(width - 2*descpadding).
		Render(lipgloss.PlaceVertical(eventHeight, lipgloss.Center, event.Name))

	from := event.Start.Format("15:04")
	to := event.End.Format("15:04")
	fromtostring := from + " - " + to
	datewidth := min(width, utf8.RuneCountInString(fromtostring))
	fromto := datestyle.
		Width(datewidth - 2*descpadding).
		Render(fromtostring)

	day := desctyle.Width(width).
		Render(getNorwegianDay(int(event.Start.Weekday())))
	desc := desctyle.Width(width - 2*descpadding).Render(event.Description)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		name,
		day,
		fromto,
		desc,
	)
}
