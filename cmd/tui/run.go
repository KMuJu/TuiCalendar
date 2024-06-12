package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kmuju/TuiCalendar/cmd/model"
	"github.com/kmuju/TuiCalendar/cmd/tui/view"
	"golang.org/x/term"
)

type teaModel struct {
	view view.BaseView
}

func initialModel(events []model.Event) teaModel {
	// events := CreateEvents()
	width, height, err := term.GetSize(0)
	if err != nil {
		return teaModel{}
	}
	// renderAmount := height/eventHeight - 1
	return teaModel{
		// NewCalendar(events, height, width, min(width/3, 40), 0, renderAmount),
		view.NewBaseView(events, width, height),
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
		default:
			m.view.HandleKey(msg.String())
			// case "h", "left":
			// 	m.week.Left()
			// case "l", "right":
			// 	m.week.Right()
		}
	}
	return m, nil
}

func (m teaModel) View() string {
	return m.view.Render()
}

func Run(events []model.Event) error {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		return err
	}
	defer f.Close()
	model := initialModel(events)
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return err
	}
	return nil
}
