package state

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/kmuju/TuiCalendar/cmd/model"
)

const (
	previewpadding = 2
	eventHeight    = 3
)

var (
	selectedColor = lipgloss.Color("#f57f52")

	daystyle = lipgloss.NewStyle().
		// Height(3).
		Bold(true).
		Background(lipgloss.Color("#3186a2"))

	eventstyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder())

	selectedstyle = lipgloss.NewStyle().
			BorderForeground(selectedColor).
			Foreground(selectedColor)

	namestyle = lipgloss.NewStyle().
			Bold(true).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			Padding(0, previewpadding)

	desctyle = lipgloss.NewStyle().
			Padding(0, previewpadding)

	datestyle = lipgloss.NewStyle().
			Padding(0, previewpadding).
			Border(lipgloss.NormalBorder(), false, false, true, false)
)

/*
Render event with format

	Name

# From - To

Renders events for the event list
*/
func renderEvent(event model.Event, width int, selected bool) string {
	style := lipgloss.NewStyle().Inherit(eventstyle)
	// wrapping := lipgloss.NewStyle().Width(width - 2)
	if selected {
		style.Inherit(selectedstyle)
	}
	from := event.Start.Format("15:04")
	to := event.End.Format("15:04")
	fromto := from + " - " + to
	return style.
		Width(width - 2).
		Render(
			lipgloss.PlaceHorizontal(width-2, lipgloss.Center,
				lipgloss.JoinVertical(
					lipgloss.Center,
					fromto,
					event.Name,
				),
			))
}

/*
Render day with format Day Date. Month
Used in the event list
*/
func renderDay(time time.Time, width, date int, month time.Month) string {
	return daystyle.
		Width(width).
		Render(
			lipgloss.Place(width, 1, lipgloss.Center, lipgloss.Center,
				getNorwegianDay(int(time.Weekday()))+" "+fmt.Sprint(date)+". "+getNorwegianMonth(month),
			),
		)
}

func getNorwegianDay(day int) string {
	switch day {
	case 0:
		return "Mandag"
	case 1:
		return "Tirsdag"
	case 2:
		return "Onsdag"
	case 3:
		return "Torsdag"
	case 4:
		return "Fredag"
	case 5:
		return "Lørdag"
	case 6:
		return "Søndag"
	}
	return "Ikke en ukedag"
}

func getNorwegianMonth(month time.Month) string {
	switch month {
	case 1:
		return "januar"
	case 2:
		return "februar"
	case 3:
		return "mars"
	case 4:
		return "april"
	case 5:
		return "mai"
	case 6:
		return "juni"
	case 7:
		return "juli"
	case 8:
		return "august"
	case 9:
		return "september"
	case 10:
		return "oktober"
	case 11:
		return "november"
	case 12:
		return "desember"
	}
	return "Finnes ikke"
}

func renderHelpText(width int, selected bool) string {
	help := []string{
		"     q - quit ",
		"ctrl-h - left ",
		"ctrl-j - down ",
		"ctrl-k - up   ",
		"ctrl-l - right",
	}
	style := lipgloss.NewStyle().
		Width(width).
		Border(lipgloss.NormalBorder(), false, false, true, false)

	if selected {
		style = style.BorderForeground(selectedColor)
	}
	return style.
		Render(
			lipgloss.Place(
				width,
				width/2,
				lipgloss.Center,
				lipgloss.Center,
				strings.Join(help, "\n"),
			),
		)
}
