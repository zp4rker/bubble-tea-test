package core

import tea "github.com/charmbracelet/bubbletea"

type Screen interface {
	View() string
	Update(tea.Msg) tea.Cmd
}
