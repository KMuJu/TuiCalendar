package state

import "github.com/kmuju/TuiCalendar/cmd/model"

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

func StartDayFilter(date int) EventFilter {
	return func(event model.Event) bool {
		_, _, d := event.Start.Date()
		return d == date
	}
}

func AfterStartDayFilter(date int) EventFilter {
	return func(event model.Event) bool {
		_, _, d := event.Start.Date()
		return d >= date
	}
}

func BetweenDaysFilter(start, end int) EventFilter {
	return func(event model.Event) bool {
		_, _, d := event.Start.Date()
		return d >= start && d < end
	}
}
