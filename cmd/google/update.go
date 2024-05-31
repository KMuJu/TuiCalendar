package google

import (
	"database/sql"

	"github.com/kmuju/TuiCalendar/cmd/db"
)

func Update(con *sql.DB) {
	events := GetEvents()
	db.Sync(con, events)
}
