package google

import (
	"database/sql"
	"fmt"

	"github.com/kmuju/TuiCalendar/cmd/db"
	"google.golang.org/api/calendar/v3"
)

func Update(service *calendar.Service, con *sql.DB) {
	events, err := GetInfo(service)
	if err != nil {
		return
	}
	db.Sync(con, createEvents(events))
}

func Delete(srv *calendar.Service, calId, eventId string) error {
	srv, err := GetService()
	if err != nil {
		return err
	}
	fmt.Printf("Deleted (%s) from (%s)\n", eventId, calId)
	err = srv.Events.Delete(calId, eventId).Do()
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}
