package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	dateHeight = 3
)

var (
	foreground = lipgloss.Color("#ebdbb2")

	dateStyle = lipgloss.NewStyle().
			Height(dateHeight).
			Bold(true).
			Align(lipgloss.Center, lipgloss.Center).
			Background(lipgloss.Color("#124294"))

	eventStyle = lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center).
			Foreground(foreground).
			BorderStyle(lipgloss.RoundedBorder()).
		// Border(eventBorder, true).
		BorderForeground(foreground)

	selectedStyle = lipgloss.NewStyle().
			Inherit(eventStyle).
			BorderLeftForeground(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#f57f52"))
)

func NewDay(date int, events []Event, isSelected bool, width, height, renderfrom, renderamount int) Day {
	from := max(0, renderfrom)
	amount := min(len(events)-1, renderfrom+renderamount)
	selected := (from + amount/2 - amount%2)
	return Day{
		Date:         date,
		Events:       events,
		NrEvents:     len(events),
		IsSelected:   isSelected,
		Selected:     selected,
		width:        width,
		height:       height,
		renderfrom:   from,
		renderamount: amount - from,
	}
}

func (day Day) Render() string {
	return day.String()
}

func (day *Day) String() string {
	doc := strings.Builder{}
	// nrEvents := len(day.Events)

	// Date
	{
		date := dateStyle.
			Width(day.width).
			Render(fmt.Sprint(day.Date))

		doc.WriteString(date + "\n")
	}

	// Events
	{
		eventWidth := day.width - 2
		for i := day.renderfrom; i < day.renderfrom+day.renderamount; i++ {
			name := renderEventDate(day.Events[i], eventWidth, day.IsSelected && i == day.Selected)
			doc.WriteString(name + "\n")
		}
	}

	return doc.String()
}

func renderEventDate(event Event, width int, selected bool) string {
	style := eventStyle
	if selected {
		style = selectedStyle
	}
	date := event.Start.Format("15:04") + " - " + event.End.Format("15:04")
	s := lipgloss.JoinVertical(lipgloss.Top, date, event.Name)
	return style.Width(width).
		// Background(c).
		// BorderBackground(c).
		Render(s)
}

func (d *Day) Up() {
	if d.Selected > 0 {
		half := d.renderamount / 2
		// log.Printf("half: %d ; selected: %d, nr: %d, renderfrom: %d\n", half, d.Selected, d.NrEvents, d.renderfrom)
		// log.Printf("%t, %t, %t\n", d.Selected > half, d.Selected < d.NrEvents-half, d.renderfrom != 0)
		if d.Selected > half && d.Selected < d.NrEvents-half && d.renderfrom != 0 {
			d.renderfrom--
		}
		d.Selected--
	}
}
func (d *Day) Down() {
	if d.Selected+1 < d.NrEvents {
		half := d.renderamount / 2
		if d.Selected+1 < d.NrEvents-half && d.Selected+1 > half && d.renderfrom+1 < d.NrEvents {
			d.renderfrom++
		}
		d.Selected++
	}
}
