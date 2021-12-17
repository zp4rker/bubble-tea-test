package core

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type Screen interface {
	View() string
	Update(tea.Msg) tea.Cmd
	Cleanup()
}

var HighlightedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("15")).
	Bold(true)

var SelectedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("11")).
	Bold(true)

func PressTo(key string, to string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	s := style.Render("Press ")
	s += lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render(key)
	s += style.Render(fmt.Sprintf(" to %v.", to))

	return s
}
