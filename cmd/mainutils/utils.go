package mainutils

import (
	"time"

	"github.com/kmuju/TuiCalendar/cmd/tui"
)

func CreateEvents() []tui.Event {
	events := []tui.Event{
		createEvent("First", "", 2024, 5, 29, 2, 30, 3, 10),
		createEvent("Abasl", "", 2024, 5, 29, 3, 30, 7, 20),
		createEvent("lascj", "", 2024, 5, 29, 9, 30, 10, 20),
		createEvent("Testing", "", 2024, 5, 29, 10, 30, 10, 40),
		createEvent("Løk", "", 2024, 5, 29, 13, 20, 11, 0),
		createEvent("HØ", "", 2024, 5, 29, 17, 34, 19, 50),
		createEvent("PO", "", 2024, 5, 29, 20, 40, 21, 0),
		createEvent("LSAH", "", 2024, 5, 29, 22, 0, 23, 25),
	}
	return events
}

func createEvent(name, description string, year, month, date, fromhour, fromminute, tohour, tominute int) tui.Event {
	return tui.Event{
		Date:        date,
		Name:        name,
		Description: description,
		Start:       time.Date(year, time.Month(month), date, fromhour, fromminute, 0, 0, time.Local),
		End:         time.Date(year, time.Month(month), date, tohour, tominute, 0, 0, time.Local),
	}
}
