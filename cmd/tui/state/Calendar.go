package state

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type Calendar struct {
	year  int
	month int

	width  int
	height int

	focus    bool
	selected int // date selected 1-31
}

var (
	daysInMonth = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

func (self *Calendar) Render() string {
	builder := strings.Builder{}
	dayFirst := getDayOfWeek(self.year, self.month, 1)
	builder.WriteString(strings.Repeat("   ", dayFirst))

	for i := 1; i <= daysInMonth[self.month]; i++ {
		year, month, date := time.Now().Date()
		s := lipgloss.NewStyle()
		if year == self.year && int(month) == self.month && date == i {
			s.Bold(true).Italic(true)
		} else if i == self.selected {
			s.Foreground(lipgloss.Color("#fb4934"))
		}

		builder.WriteString(s.Render(lipgloss.PlaceHorizontal(3, lipgloss.Right, fmt.Sprint(i))))

		if getDayOfWeek(self.year, self.month, i) == 6 /* && i != daysInMonth[self.month]  */ {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}

func NewCalendar(t time.Time, width, height int) Calendar {
	year, month, _ := t.Date()
	if isLeapYear(year) {
		daysInMonth[2] = 29
	}
	return Calendar{
		year:     year,
		month:    int(month),
		width:    width,
		height:   height,
		selected: 4,
	}
}

func isLeapYear(year int) bool {
	if (year%4 == 0) && (year%100 != 0) {
		return true
	}
	if year%400 == 0 {
		return true
	}
	return false
}

func getDayOfWeek(year, month, day int) int {
	a := (14 - month) / 12
	y := year - a
	m := month + (12 * a) - 2
	d := (day + y + (y / 4) - (y / 100) + (y / 400) + ((31 * m) / 12)) % 7
	return d
}
