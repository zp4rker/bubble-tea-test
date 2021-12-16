package mainmodel

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/core"
)

type Model struct {
	currentScreen core.Screen
}

var MainModel = Model{}

func (m *Model) LoadScreen(screen core.Screen) {
	m.currentScreen = screen
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (Model) View() string {
	return MainModel.currentScreen.View()
}

func (Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return MainModel, MainModel.currentScreen.Update(msg)
}
