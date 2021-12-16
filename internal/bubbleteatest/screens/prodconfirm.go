package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubbleteatest/internal/bubbleteatest/core"
	"github.com/zp4rker/bubbleteatest/internal/bubbleteatest/mainmodel"
)

type screen_prodconfirm struct {
	product string

	choices  []string
	cursor   int
	selected int
}

func ProdConfirm(product string) core.Screen {
	return &screen_prodconfirm{
		product:  product,
		choices:  []string{"Yes", "No"},
		selected: -1,
	}
}

func (scr *screen_prodconfirm) View() string {
	s := fmt.Sprintf("You chose '%v'. Is that what you would like to order?\n\n", scr.product)

	for i, choice := range scr.choices {
		cursor := " "
		if scr.cursor == i {
			cursor = ">"
		}

		checked := " "
		if scr.selected == i {
			checked = "*"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
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

		case " ":
			if scr.selected != scr.cursor {
				scr.selected = scr.cursor
			}

		case "enter":
			if scr.selected != -1 {
				mainmodel.MainModel.LoadScreen(Exit(scr.product, scr.choices[scr.selected] == "Yes"))
			}
		}
	}

	return nil
}
