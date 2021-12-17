package marketapp

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/core"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/mainmodel"
)

type screen_prodconfirm struct {
	product string

	choices []string
	cursor  int
}

func ProdConfirm(product string) core.Screen {
	return &screen_prodconfirm{
		product: product,
		choices: []string{"Yes", "No"},
	}
}

func (scr *screen_prodconfirm) View() string {
	s := fmt.Sprintf("\nYou chose '%v'. Is that what you would like to order?\n\n", scr.product)

	for i, choice := range scr.choices {
		if scr.cursor == i {
			s += core.HighlightedStyle.Render(fmt.Sprintf("\r> %s\n", choice))
		} else {
			s += fmt.Sprintf("\r  %s\n", choice)
		}
	}

	s += fmt.Sprintf("\n%v\n", core.PressTo("q", "quit"))

	return s
}

func (scr *screen_prodconfirm) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			mainmodel.MainModel.LoadMainScreen()

		case "up":
			if scr.cursor > 0 {
				scr.cursor--
			}

		case "down":
			if scr.cursor < len(scr.choices)-1 {
				scr.cursor++
			}

		case " ", "enter":
			mainmodel.MainModel.LoadScreen(Exit(scr.product, scr.choices[scr.cursor] == "Yes"))
		}
	}

	return nil
}

func (scr *screen_prodconfirm) Cleanup() {
	scr.cursor = 0
}
