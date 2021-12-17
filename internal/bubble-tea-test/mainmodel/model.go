package mainmodel

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/core"
)

type Model struct {
	mainScreen    core.Screen
	currentScreen core.Screen
}

var MainModel = Model{}

func (m *Model) LoadScreen(screen core.Screen) {
	if m.currentScreen != nil {
		defer func(screen core.Screen) {
			screen.Cleanup()
		}(m.currentScreen)
	}

	m.currentScreen = screen

	if m.mainScreen == nil {
		m.mainScreen = screen
	}
}

func (m *Model) LoadMainScreen() {
	m.currentScreen = m.mainScreen
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
