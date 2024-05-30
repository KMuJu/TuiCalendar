package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

const (
	descpadding = 2
)

var (
	daystyle = lipgloss.NewStyle().
		// Height(3).
		Bold(true).
		Background(lipgloss.Color("#3186a2"))

	eventstyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder())

	selectedstyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.Color("#f57f52")).
			Foreground(lipgloss.Color("#f57f52"))

	namestyle = lipgloss.NewStyle().
			Bold(true).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			Padding(0, descpadding)

	desctyle = lipgloss.NewStyle().
			Padding(0, descpadding)

	datestyle = lipgloss.NewStyle().
			Padding(0, descpadding).
			Border(lipgloss.NormalBorder(), false, false, true, false)
)

func NewCalendar(events []Event, height, width, listWidth, renderFrom, renderAmount int) Calendar {
	selected := renderFrom + renderAmount/2
	return Calendar{
		events:       events,
		height:       height,
		width:        width,
		listWidth:    listWidth,
		renderFrom:   renderFrom,
		renderAmount: renderAmount,
		selected:     selected,
	}
}

func (c *Calendar) Render() string {
	listDoc := strings.Builder{}
	descDoc := strings.Builder{}

	// List
	{
		width := c.listWidth
		lastDate := 0
		lastMonth := 0
		for i := c.renderFrom; i < c.renderFrom+c.renderAmount; i++ {
			event := c.events[i]
			_, month, date := event.Start.Date()
			if lastMonth != int(month) || lastDate < date {
				listDoc.WriteString(renderDay(event.Start, width, date, month) + "\n")
			}
			lastDate = date
			lastMonth = int(month)

			// Render event
			listDoc.WriteString(renderEvent(event, width, i == c.selected) + "\n")
		}

		{
			event := c.events[c.selected]
			descDoc.WriteString(renderDescription(event, c.width-c.listWidth))
		}

	}

	listString := listDoc.String()
	descriptionString := descDoc.String()
	middle := strings.Repeat(" ┃\n", len(strings.Split(listString, "\n")))
	return lipgloss.JoinHorizontal(lipgloss.Top, listString, middle, descriptionString)
}

func (c *Calendar) String() string {
	return c.Render()
}

func (c *Calendar) Up() {
	if c.selected > 0 {
		if c.selected == c.renderFrom {
			c.renderFrom--
		}
		c.selected--
	}
}
func (c *Calendar) Down() {
	if c.selected+1 < len(c.events) {
		if c.selected == c.renderFrom+c.renderAmount-1 {
			c.renderFrom++
		}
		c.selected++
	}
}

func renderDay(time time.Time, width, date int, month time.Month) string {
	return daystyle.
		Render(
			lipgloss.Place(width, 1, lipgloss.Center, lipgloss.Center,
				getNorwegianDay(int(time.Weekday()))+" "+fmt.Sprint(date)+". "+getNorwegianMonth(month),
			),
		)
}

func renderEvent(event Event, width int, selected bool) string {
	style := lipgloss.NewStyle().Inherit(eventstyle)
	if selected {
		style.Inherit(selectedstyle)
	}
	from := event.Start.Format("15:04")
	to := event.End.Format("15:04")
	fromto := from + " - " + to
	return style.
		Render(
			lipgloss.PlaceHorizontal(width-2, lipgloss.Center,
				lipgloss.JoinVertical(lipgloss.Center, fromto, event.Name),
			))
}

func renderDescription(event Event, width int) string {
	name := namestyle.
		Width(width).
		Render(lipgloss.PlaceVertical(3, lipgloss.Center, event.Name))
		// from :=
	from := event.Start.Format("15:04")
	to := event.End.Format("15:04")
	fromto := datestyle.Render(from + " - " + to)
	day := desctyle.
		Render(getNorwegianDay(int(event.Start.Weekday())))
	desc := desctyle.Width(width).Render(event.Description)
	return lipgloss.JoinVertical(lipgloss.Left, name, day, fromto, desc)
}

func truncateString(input string, maxlen int) string {
	if maxlen < 3 {
		return "..."
	}
	if len(input) > maxlen {
		return input[:maxlen-3] + "..."
	}
	return input
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
