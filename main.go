package main

import (
	"context"
	"fmt"
	"os"
	"slices"

	"github.com/kmuju/TuiCalendar/cmd/db"
	"github.com/kmuju/TuiCalendar/cmd/google"
	"github.com/kmuju/TuiCalendar/cmd/model"
	"github.com/kmuju/TuiCalendar/cmd/tui"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name: "TuiCalendar",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "sync"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Bool("sync") {
				// _, err := exec.Command("./sync").Output()
				// return err
				update()
				return nil
			}
			con, err := db.InitDB()
			if err != nil {
				return err
			}
			events, err := db.GetEvents(con)
			if err != nil {
				return err
			}
			slices.SortFunc(events, func(a, b model.Event) int {
				if a.Start.Before(b.Start) {
					return -1
				}
				if a.Start.After(b.Start) {
					return 1
				}
				return 0
			})
			tui.Run(events)
			return nil
		},
	}

	cmd.Run(context.Background(), os.Args)
}

func update() {

	con, err := db.InitDB()
	if err != nil {
		return
	}
	google.Update(con)
	dbevents, err := db.GetEvents(con)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("New events:\n")
	for _, ev := range dbevents {
		fmt.Printf("%v\n", ev)
	}
}
