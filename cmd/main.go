package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kmuju/TuiCalendar/cmd/mainutils"
	"github.com/kmuju/TuiCalendar/cmd/tui"
)

type model struct {
	day tui.Day
}

func initialModel() model {
	events := mainutils.CreateEvents()
	day := tui.NewDay(
		12,
		events,
		true,
		40, 20,
		2, 5,
	)
	return model{
		day: day,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k", "up":
			m.day.Up()
		case "j", "down":
			m.day.Down()
			// case "h", "left":
			// 	if m.cal.X > 0 {
			// 		m.cal.X--
			// 	}
			// case "l", "right":
			// 	if m.cal.X+1 < m.cal.Width {
			// 		m.cal.X++
			// 	}
		}
	}
	return m, nil
}

func (m model) View() string {
	return m.day.Render()
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
