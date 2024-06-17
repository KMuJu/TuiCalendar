package state

import (
	"time"

	"github.com/kmuju/TuiCalendar/cmd/model"
)

type EventController struct {
	Events []model.Event
}

type EventFilter func(event model.Event) bool

func (self *EventController) GetAllEvents() []model.Event {
	return self.Events
}

func (self *EventController) GetEvents(filter EventFilter) []model.Event {
	newevents := []model.Event{}

	for _, event := range self.Events {
		if filter(event) {
			newevents = append(newevents, event)
		}
	}

	return newevents
}

func StartDayFilter(year, month, date int) EventFilter {
	return func(event model.Event) bool {
		y, m, d := event.Start.Date()
		return d == date && y == year && int(m) == month
	}
}

func AfterDayFilter(year, month, date int) EventFilter {
	return func(event model.Event) bool {
		y, m, d := event.Start.Date()
		return d >= date && y >= year && int(m) >= month
	}
}

func BetweenTimesFilter(start, end time.Time) EventFilter {
	return func(event model.Event) bool {
		return event.Start.After(start) && event.End.Before(end)
	}
}
