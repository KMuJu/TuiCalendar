package state

import (
	"time"
)

type Sidebar struct {
	calendar Calendar
	width    int
	height   int
	focus    bool
}

func NewSidebar(width, height int, daysWithEvent []int) *Sidebar {
	return &Sidebar{
		calendar: NewCalendar(time.Now(), width, height, daysWithEvent),
		width:    width,
		height:   height,
		focus:    false,
	}
}

func (self *Sidebar) Focus()             { self.focus = true }
func (self *Sidebar) FocusLost()         { self.focus = false }
func (_ *Sidebar) InnerFocus(_ int) bool { return false }
func (self *Sidebar) HandleKey(key string) {
	switch key {
	case "k", "up":
		self.calendar.up()
	case "j", "down":
		self.calendar.down()
	case "l", "right":
		self.calendar.right()
	case "h", "left":
		self.calendar.left()
	}
}

func (self *Sidebar) GetSelected() int {
	return self.calendar.GetSelected()
}

func (self *Sidebar) Render() string {
	help := renderHelpText(self.width, self.focus)
	cal := self.calendar.Render()
	return help + "\n" + cal
}
