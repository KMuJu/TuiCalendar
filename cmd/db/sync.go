package db

import (
	"database/sql"
	"fmt"
	"slices"

	"github.com/kmuju/TuiCalendar/cmd/model"
)

func Sync(db *sql.DB, newEvents []model.Event) error {
	oldEvents, err := GetEvents(db)
	if err != nil {
		return err
	}
	fmt.Printf("Sync\n%v\n", newEvents)
	for _, e := range newEvents {
		if slices.Contains(oldEvents, e) {
			continue
		}
		fmt.Printf("Event is synced %v\n", e)
		err = InsertEvent(db, e)
		if err != nil {
			fmt.Printf("Could not insert event: %v\n", e)
		}
	}
	return nil
}
