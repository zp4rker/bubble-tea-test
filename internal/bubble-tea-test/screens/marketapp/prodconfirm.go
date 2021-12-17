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
		cursor := " "
		if scr.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func (scr *screen_prodconfirm) Update(msg tea.Msg) tea.Cmd {
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
			mainmodel.MainModel.LoadScreen(Exit(scr.product, scr.choices[scr.cursor] == "Yes"))
		}
	}

	return nil
}
