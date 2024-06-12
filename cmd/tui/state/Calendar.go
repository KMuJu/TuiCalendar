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
	col      int
	row      int
}

var (
	daysInMonth = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	cal         = [][]int{}
)

func (self *Calendar) up() {
	if self.row != 0 {
		self.row--
	}
	selected := cal[self.row][self.col]
	if selected == 0 {
		// Find the first element in the row
		for i := 0; i < 7; i++ {
			if cal[self.row][i] != 0 {
				selected = cal[self.row][i]
				break
			}
		}
	}
	self.selected = selected
}

func (self *Calendar) down() {
	if self.row+1 != len(cal) {
		self.row++
	}
	selected := cal[self.row][self.col]
	if selected == 0 {
		// Find the first element in the row
		for i := 6; i <= 0; i-- {
			if cal[self.row][i] != 0 {
				selected = cal[self.row][i]
				break
			}
		}
	}
	self.selected = selected
}

func (self *Calendar) left() {
	if self.col != 0 {
		if cal[self.row][self.col-1] != 0 {
			self.col--
			self.selected = cal[self.row][self.col]
		}
	}
}

func (self *Calendar) right() {
	if self.col != 6 {
		if cal[self.row][self.col+1] != 0 {
			self.col++
			self.selected = cal[self.row][self.col]
		}
	}
}

func (self *Calendar) Render() string {
	builder := strings.Builder{}
	dayFirst := getDayOfWeek(self.year, self.month, 1)
	builder.WriteString(strings.Repeat("   ", dayFirst))

	for i := 1; i <= daysInMonth[self.month]; i++ {
		year, month, date := time.Now().Date()
		s := lipgloss.NewStyle()
		if year == self.year && int(month) == self.month && date == i {
			s = s.Bold(true).Italic(true).Underline(true)
			// s = s
		}
		if i == self.selected {
			s = s.Foreground(lipgloss.Color("#fb4934"))
		}

		builder.WriteString(s.Render(lipgloss.PlaceHorizontal(3, lipgloss.Right, fmt.Sprint(i))))

		if getDayOfWeek(self.year, self.month, i) == 6 /* && i != daysInMonth[self.month]  */ {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}

func NewCalendar(t time.Time, width, height int) Calendar {
	year, m, _ := t.Date()
	month := int(m)
	if isLeapYear(year) {
		daysInMonth[2] = 29
	}

	calcCalendar(year, month)

	return Calendar{
		year:     year,
		month:    month,
		width:    width,
		height:   height,
		selected: 1,

		row: 0,
		col: getDayOfWeek(year, month, 1),
	}
}

func calcCalendar(year, month int) {
	// dayFirst := getDayOfWeek(year, month, 1)
	row := 0

	cal = make([][]int, 1)
	cal[0] = make([]int, 7)

	// for i := 0; i < dayFirst; i++ {
	// 	cal[0][i] = 0
	// }

	for i := 1; i <= daysInMonth[month]; i++ {
		day := getDayOfWeek(year, month, i)
		cal[row][day] = i

		if day == 6 {
			row++
			cal = append(cal, make([]int, 7))
		}
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
