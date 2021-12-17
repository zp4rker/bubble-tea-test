package appselect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/core"
)

type screen_comingsoon struct {
	app string
}

func ComingSoon(app string) core.Screen {
	return &screen_comingsoon{
		app: app,
	}
}

func (scr *screen_comingsoon) View() string {
	s := fmt.Sprintf("\n%s is coming soon!\n\n", scr.app)

	s += "\nPress any key to quit.\n"

	return s
}

func (scr *screen_comingsoon) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case tea.KeyMsg:
		return tea.Quit
	}

	return nil
}
