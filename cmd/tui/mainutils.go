package tui

import (
	"time"

	"github.com/kmuju/TuiCalendar/cmd/model"
)

func CreateEvents() []model.Event {
	events := []model.Event{
		createEvent("a", "First", "Dette er et event",
			2024, 5, 15,
			2, 30,
			3, 10),
		createEvent("b", "Abasl", "YAYAY",
			2024, 5, 15,
			3, 30,
			7, 20),
		createEvent("c", "lascj", "Paødslfk",
			2024, 5, 19,
			9, 30,
			10, 20),
		createEvent("d", "Testing", "øfa",
			2024, 5, 20,
			10, 30,
			10, 40),
		createEvent("e", "Løk", "Aleks er en løk",
			2024, 5, 25,
			13, 20,
			11, 0),
		createEvent("s", "HØ", "",
			2024, 5, 25,
			17, 34,
			19, 50),
		createEvent("w", "PO", "",
			2024, 5, 27,
			20, 40,
			21, 0),
		createEvent("j", "LSAH", "",
			2024, 5, 29,
			22, 0,
			23, 25),
	}
	return events
}

func createEvent(Id, name, description string, year, month, date, fromhour, fromminute, tohour, tominute int) model.Event {
	return model.Event{
		Id:          Id,
		Date:        date,
		Name:        name,
		Description: description,
		Start:       time.Date(year, time.Month(month), date, fromhour, fromminute, 0, 0, time.Local),
		End:         time.Date(year, time.Month(month), date, tohour, tominute, 0, 0, time.Local),
	}
}
