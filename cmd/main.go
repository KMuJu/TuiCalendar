package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kmuju/TuiCalendar/cmd/tui"
)

type model struct {
	day  tui.Day
	week tui.Week
}

func initialModel() model {
	events := tui.CreateEvents()
	day := tui.NewDay(
		12,
		events,
		true,
		40, 20,
		2, 5,
	)
	week := tui.NewWeek(108, 0)
	return model{
		day:  day,
		week: week,
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
			m.week.Up()
		case "j", "down":
			m.week.Down()
		case "h", "left":
			m.week.Left()
		case "l", "right":
			m.week.Right()
		}
	}
	return m, nil
}

func (m model) View() string {
	return m.week.Render()
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
