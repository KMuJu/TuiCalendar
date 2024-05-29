package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
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
			Border(lipgloss.NormalBorder(), false, false, true, false)

	desctyle = lipgloss.NewStyle()
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
				listDoc.WriteString(renderDay(event.Start, width) + "\n")
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
	middle := strings.Repeat("┃\n", len(strings.Split(listString, "\n")))
	return lipgloss.JoinHorizontal(lipgloss.Top, listString, middle, descriptionString)
}

func (c *Calendar) String() string {
	return c.Render()
}

func (c *Calendar) Up() {
	if c.selected > 0 {
		c.selected--
	}
}
func (c *Calendar) Down() {
	if c.selected+1 < len(c.events) {
		c.selected++
	}
}

func renderDay(time time.Time, width int) string {
	return daystyle.
		Render(
			lipgloss.Place(width, 1, lipgloss.Center, lipgloss.Center,
				getNorwegianDay(int(time.Weekday()))),
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
		Render(lipgloss.PlaceHorizontal(width, lipgloss.Center, event.Name))
	desc := desctyle.Width(width).Render(event.Description)
	return lipgloss.JoinVertical(lipgloss.Left, name, desc)
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
