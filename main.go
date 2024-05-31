package main

import (
	"fmt"

	d "github.com/kmuju/TuiCalendar/cmd/db"
	"github.com/kmuju/TuiCalendar/cmd/model"
	"github.com/kmuju/TuiCalendar/cmd/tui"
)

func main() {
	db, err := d.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	e := model.Event{
		Name:        "Test",
		Description: "ALS",
	}
	d.InsertEvent(db, e)

	dbevents, err := d.GetEvents(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ev := range dbevents {
		fmt.Printf("%v\n", ev)
	}

	events := tui.CreateEvents()
	d.Sync(db, events)

	dbevents, err = d.GetEvents(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("New events:\n")
	for _, ev := range dbevents {
		fmt.Printf("%v\n", ev)
	}
}
