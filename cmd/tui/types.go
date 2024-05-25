package tui

import (
	"time"

	"github.com/charmbracelet/lipgloss/table"
)

type Day struct {
	Events       []Event
	NrEvents     int
	Date         int // 1-31
	Month        int // 1-12
	width        int
	height       int
	IsSelected   bool
	Selected     int
	renderfrom   int
	renderamount int
}

type WeekTable struct {
	Days      [7]Day
	Nr        int // 1-52
	Width     int
	Height    int
	CellWidth int
	Table     table.Table
}

type Event struct {
	Date        int
	Description string
	Name        string
	Start       time.Time
	End         time.Time
}
