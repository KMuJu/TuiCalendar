package state

import (
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	"github.com/kmuju/TuiCalendar/cmd/model"
)

type EventPreview struct {
	width  int
	height int
	focus  bool
}

func NewPreviewer(width, height int) *EventPreview {
	return &EventPreview{width: width, height: height, focus: false}
}

func (self *EventPreview) Focus()             { self.focus = true }
func (self *EventPreview) FocusLost()         { self.focus = false }
func (_ *EventPreview) InnerFocus(_ int) bool { return false }
func (_ *EventPreview) HandleKey(_ string)    {}

func (self *EventPreview) HandleWidthChange(delta int) {
	if self.width+delta > 0 {
		self.width += delta
	}
}

func (self *EventPreview) HandleHeightChange(delta int) {
	if self.height+delta > 0 {
		self.height += delta
	}
}

func (self *EventPreview) Render(event model.Event) string {
	namestyle := lipgloss.NewStyle().Inherit(namestyle)
	datestyle := lipgloss.NewStyle().Inherit(datestyle)
	namestyle = namestyle.Padding(0, previewpadding)
	datestyle = datestyle.Padding(0, previewpadding)

	if self.focus {
		namestyle.BorderForeground(selectedColor)
		datestyle.BorderForeground(selectedColor)
	}

	width := self.width
	name := namestyle.
		Width(width - 2*previewpadding).
		Render(lipgloss.PlaceVertical(eventHeight, lipgloss.Center, event.Name))

	from := event.Start.Format("15:04")
	to := event.End.Format("15:04")
	fromtostring := from + " - " + to
	datewidth := min(width-2*previewpadding, utf8.RuneCountInString(fromtostring)+2*previewpadding)
	fromto := datestyle.
		Width(datewidth).
		Render(fromtostring)

	day := desctyle.Width(width).
		Render(getNorwegianDay(int(event.Start.Weekday())))
	desc := desctyle.Width(width - 2*previewpadding).Render(event.Description)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		name,
		day,
		fromto,
		desc,
	)
}
