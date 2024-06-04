package db

import (
	"database/sql"
	"fmt"
	"slices"
	"time"

	"github.com/kmuju/TuiCalendar/cmd/model"
)

func Sync(con *sql.DB, newEvents []model.Event) error {
	oldEvents, err := GetEvents(con)
	if err != nil {
		return err
	}
	fmt.Printf("Sync\n%v\n", newEvents)
	for _, e := range newEvents {
		if slices.Contains(oldEvents, e) {
			continue
		}
		if e.Status == "cancelled" {
			DeleteEvent(con, e.Id)
			fmt.Printf("Deleting event: %+v\n", e)
			continue
		}
		err = InsertEvent(con, e)
		if err != nil {
			fmt.Printf("Could not insert event: %v\n", e)
			fmt.Println(err)
		}
	}
	now := time.Now()
	for _, e := range oldEvents {
		if slices.Contains(newEvents, e) {
			continue
		}
		if e.Start.Before(now) {
			continue
		}
		err = DeleteEvent(con, e.Id)
		if err != nil {
			fmt.Printf("Could not delete:\n%s\n", err.Error())
		}
	}
	return nil
}
