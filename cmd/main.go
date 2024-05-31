package main

import (
	"github.com/kmuju/TuiCalendar/cmd/db"
	"github.com/kmuju/TuiCalendar/cmd/tui"
)

func main() {
	con, err := db.InitDB()
	if err != nil {
		return
	}
	events, err := db.GetEvents(con)
	if err != nil {
		return
	}
	tui.Run(events)
}
