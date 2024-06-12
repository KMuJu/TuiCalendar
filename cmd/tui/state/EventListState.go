package state

import (
	"strings"

	"github.com/kmuju/TuiCalendar/cmd/model"
)

type EventListState struct {
	events []model.Event
	length int

	width    int
	height   int
	start    int
	focus    bool
	selected int
}

func NewEventList(events []model.Event, width, height, start int, focus bool) EventListState {
	return EventListState{
		events:   events,
		length:   len(events),
		width:    width,
		height:   height,
		start:    start,
		selected: start,
		focus:    focus,
	}
}

func (self *EventListState) Focus()     { self.focus = true }
func (self *EventListState) FocusLost() { self.focus = false }

func (self *EventListState) Render() string {
	listDoc := strings.Builder{}

	width := self.width
	lastDate := 0
	lastMonth := 0
	height := 0
	for i := self.start; i < len(self.events); i++ {
		if height+3 > self.height { // If the next event is outside height
			break
		}
		event := self.events[i]
		_, month, date := event.Start.Date()
		if lastMonth != int(month) || lastDate < date {
			listDoc.WriteString(renderDay(event.Start, width, date, month) + "\n")
			height++
		}
		lastDate = date
		lastMonth = int(month)

		// Render event
		listDoc.WriteString(renderEvent(event, width, i == self.selected) + "\n")
		height += 3
	}
	return listDoc.String()
}

func (self *EventListState) HandleWidthChange(delta int) {
	if self.width+delta > 0 {
		self.width += delta
	}
}

func (self *EventListState) HandleHeightChange(delta int) {
	if self.height+delta > 0 {
		self.height += delta
	}
}

func (self *EventListState) Len() int {
	return self.length
}

/*Moves the start by delta events*/
func (self *EventListState) MoveStart(delta int) {
	if self.start+delta < self.length && self.start+delta >= 0 {
		self.start += delta
	}
}

func (self *EventListState) MoveSelected(delta int) {
	if self.selected+delta < self.length && self.selected+delta >= 0 {
		self.selected += delta
	}
}

func (self *EventListState) Up() {
	self.MoveSelected(-1)
}

func (self *EventListState) Down() {
	self.MoveSelected(1)
}

func (self *EventListState) GetSelectedEvent() model.Event {
	return self.events[self.selected]
}

func (self *EventListState) HandleKey(input string) {
	switch input {
	case "k", "up":
		self.Up()
	case "j", "down":
		self.Down()
	}
}
