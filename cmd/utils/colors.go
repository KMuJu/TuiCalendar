package utils

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/crazy3lf/colorconv"
)

func CreateColors(length int) []lipgloss.Color {
	colors := make([]lipgloss.Color, length)
	hue := 0
	for i := 0; i < length; i++ {
		c, _ := colorconv.HSLToColor(float64(hue), 0.5, 0.3)
		correct, _ := strings.CutPrefix(colorconv.ColorToHex(c), "0x")
		colors[i] = lipgloss.Color("#" + correct)
		hue = (hue + 40) % 180
	}
	return colors
}
