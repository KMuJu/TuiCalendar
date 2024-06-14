package state

import (
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var (
	selectedday = false
	today       = false
)

func Reset(controller *EventController, list *EventList) {
	events := controller.GetAllEvents()
	list.SetEvents(events)
}

func ToggleTodayFilter(controller *EventController, list *EventList) {
	if selectedday || today { // if another filter is active reseet
		Reset(controller, list)
	} else {
		date := time.Now().Day()
		list.SetEvents(
			controller.GetEvents(
				StartDayFilter(date),
			),
		)
	}
	today = !today
}

func ToggleSelectedDay(controller *EventController, list *EventList, day int) {
	if selectedday || today {
		Reset(controller, list)
	} else {
		list.SetEvents(
			controller.GetEvents(
				StartDayFilter(day),
			),
		)
	}
	selectedday = !selectedday
}

func Render(controller *EventController, list *EventList, width int) string {
	filters := []string{
		"Today: ",
		"Selected: ",
	}
	if today {
		filters[0] += ""
		log.Printf("Today filter: %s\n", filters[0])
	}
	if selectedday {
		filters[1] += ""
		log.Printf("Selected filter: %s\n", filters[0])
	}
	return lipgloss.NewStyle().
		Width(width).
		Render(
			lipgloss.JoinVertical(lipgloss.Left,
				"Filters:\n────────",
				lipgloss.PlaceHorizontal(width, lipgloss.Center, strings.Join(filters, "\n")),
			),
		)
}
