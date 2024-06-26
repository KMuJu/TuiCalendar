package state

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kmuju/TuiCalendar/cmd/model"
)

type EventList struct {
	events []model.Event
	length int

	width    int
	height   int
	start    int
	focus    bool
	selected int
}

func NewEventList(events []model.Event, width, height, start int, focus bool) *EventList {
	return &EventList{
		events:   events,
		length:   len(events),
		width:    width,
		height:   height,
		start:    start,
		selected: start,
		focus:    focus,
	}
}

func (self *EventList) Focus()             { self.focus = true }
func (self *EventList) FocusLost()         { self.focus = false }
func (_ *EventList) InnerFocus(_ int) bool { return false }
func (self *EventList) SetEvents(events []model.Event) {
	self.events = events
	self.length = len(events)
	self.selected = min(self.selected, self.length-1)
}

func (self *EventList) Render() string {
	if self.length == 0 {
		return lipgloss.NewStyle().Width(self.width).Render("")
	}
	listDoc := strings.Builder{}

	width := self.width
	lastDate := 0
	lastMonth := 0
	height := 0
	for i := self.start; i < self.length; i++ {
		if height+3 > self.height { // If the next event is outside height
			break
		}
		event := self.events[i]
		_, month, date := event.Start.Date()
		// log.Printf("Prev: %d %d ; Event %d %d\n", lastMonth, lastDate, int(month), date)
		if lastMonth != int(month) || lastDate < date {
			// log.Printf("Rendered day %d %d\n", int(month), date)
			listDoc.WriteString(renderDay(event.Start, width, date, month) + "\n")
			height++
		}
		lastDate = date
		lastMonth = int(month)

		// Render event
		listDoc.WriteString(renderEvent(event, width, i == self.selected && self.focus) + "\n")
		height += 3
	}
	return listDoc.String()
}

func (self *EventList) HandleWidthChange(delta int) {
	if self.width+delta > 0 {
		self.width += delta
	}
}

func (self *EventList) HandleHeightChange(delta int) {
	if self.height+delta > 0 {
		self.height += delta
	}
}

func (self *EventList) Len() int {
	return self.length
}

/*Moves the start by delta events*/
func (self *EventList) MoveStart(delta int) {
	if self.start+delta < self.length && self.start+delta >= 0 {
		self.start += delta
	}
}

func (self *EventList) MoveSelected(delta int) {
	if self.selected+delta < self.length && self.selected+delta >= 0 {
		self.selected += delta
	}
}

func (self *EventList) Up() {
	self.MoveSelected(-1)
}

func (self *EventList) Down() {
	self.MoveSelected(1)
}

func (self *EventList) GetSelectedEvent() model.Event {
	if self.selected >= self.length || self.selected < 0 {
		return model.Event{}
	}
	return self.events[self.selected]
}

func (self *EventList) HandleKey(input string) {
	switch input {
	case "k", "up":
		self.Up()
	case "j", "down":
		self.Down()
	}
}
