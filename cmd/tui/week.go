package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func NewWeek(width, height int) Week {
	var days [7]Day
	for i := 0; i < 7; i++ {
		events := CreateEvents()
		days[i] = NewDay(i+1, events,
			i == 0,
			width/7, 20,
			0, 5)
	}
	week := Week{
		Days:     days,
		Nr:       23,
		width:    width,
		height:   height,
		selected: 0,
	}
	return week
}

func (w *Week) Render() string {
	renders := make([]string, 7)
	for i := 0; i < 7; i++ {
		renders[i] = w.Days[i].Render()
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, renders...)
}

func (week *Week) Update() {

}

func (w *Week) Up() {
	for i := 0; i < 7; i++ {
		w.Days[i].Up()
	}
}
func (w *Week) Down() {
	for i := 0; i < 7; i++ {
		w.Days[i].Down()
	}
}

func (w *Week) Left() {
	if w.selected != 0 {
		w.Days[w.selected].IsSelected = false
		w.selected--
		w.Days[w.selected].IsSelected = true
	}
}
func (w *Week) Right() {
	if w.selected != 6 {
		w.Days[w.selected].IsSelected = false
		w.selected++
		w.Days[w.selected].IsSelected = true
	}
}
