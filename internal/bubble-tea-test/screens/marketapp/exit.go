package marketapp

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/core"
)

type screen_exit struct {
	product   string
	purchased bool
}

func Exit(product string, purchased bool) core.Screen {
	return &screen_exit{
		product:   product,
		purchased: purchased,
	}
}

func (scr *screen_exit) View() string {
	s := "\nYou "
	if scr.purchased {
		s += fmt.Sprintf("purchased %v.\n", scr.product)
	} else {
		s += "did not purchase anything.\n"
	}

	s += "\nPress any key to quit.\n"

	return s
}

func (*screen_exit) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case tea.KeyMsg:
		return tea.Quit
	}

	return nil
}
