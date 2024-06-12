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

func NewSidebar(width, height int) *Sidebar {
	return &Sidebar{
		calendar: NewCalendar(time.Now(), width, height),
		width:    width,
		height:   height,
		focus:    false,
	}
}

func (self *Sidebar) Focus()                { self.focus = true }
func (self *Sidebar) FocusLost()            { self.focus = false }
func (self *Sidebar) InnerFocus(_ int) bool { return false }

func (self *Sidebar) Render() string {
	help := renderHelpText(self.width, self.focus)
	cal := self.calendar.Render()
	return help + "\n" + cal
}
