package appselect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/core"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/mainmodel"
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
	s := fmt.Sprintf("\n%s is coming soon!\n", scr.app)

	s += fmt.Sprintf("\n%v\n", core.PressTo("any key", "quit"))

	return s
}

func (scr *screen_comingsoon) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case tea.KeyMsg:
		mainmodel.MainModel.LoadMainScreen()
	}

	return nil
}

func (scr *screen_comingsoon) Cleanup() {
	// Nothing to cleanup
}
