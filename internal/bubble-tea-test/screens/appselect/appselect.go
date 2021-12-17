package appselect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/core"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/mainmodel"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/screens/marketapp"
)

type screen_appselect struct {
	choices []string
	apps    map[string]core.Screen
	cursor  int
}

func AppSelect() core.Screen {
	return &screen_appselect{
		choices: []string{
			"Market App",
			"Simple Game",
		},
		apps: map[string]core.Screen{
			"Market App":  marketapp.ProdChoose(),
			"Simple Game": ComingSoon("Simple Game"),
		},
	}
}

func (scr *screen_appselect) View() string {
	s := "\nChoose an app:\n\n"

	for i, app := range scr.choices {
		if scr.cursor == i {
			s += core.HighlightedStyle.Render(fmt.Sprintf("\r> %s\n", app))
		} else {
			s += fmt.Sprintf("\r  %s\n", app)
		}
	}

	s += fmt.Sprintf("\n%v\n", core.PressTo("q", "quit"))

	return s
}

func (scr *screen_appselect) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return tea.Quit

		case "up":
			if scr.cursor > 0 {
				scr.cursor--
			}

		case "down":
			if scr.cursor < len(scr.choices)-1 {
				scr.cursor++
			}

		case " ", "enter":
			mainmodel.MainModel.LoadScreen(scr.apps[scr.choices[scr.cursor]])
		}
	}

	return nil
}

func (scr *screen_appselect) Cleanup() {
	scr.cursor = 0
}
