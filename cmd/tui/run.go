package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kmuju/TuiCalendar/cmd/model"
)

type teaModel struct {
	cal Calendar
}

func initialModel(events []model.Event) teaModel {
	// events := CreateEvents()
	return teaModel{
		NewCalendar(events, 20, 100, 40, 0, 5),
	}
}

func (m teaModel) Init() tea.Cmd {
	return nil
}

func (m teaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k", "up":
			m.cal.Up()
		case "j", "down":
			m.cal.Down()
			// case "h", "left":
			// 	m.week.Left()
			// case "l", "right":
			// 	m.week.Right()
		}
	}
	return m, nil
}

func (m teaModel) View() string {
	return m.cal.Render()
}

func Run(events []model.Event) error {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		return err
	}
	defer f.Close()
	p := tea.NewProgram(initialModel(events), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return err
	}
	return nil
}
