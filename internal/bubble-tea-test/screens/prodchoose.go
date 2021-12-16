package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/core"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/mainmodel"
)

type screen_prodchoose struct {
	choices  []string
	cursor   int
	selected int
}

func ProdChoose() core.Screen {
	return &screen_prodchoose{
		choices:  []string{"Carrots", "Cucumbers", "Apples", "Bananas"},
		selected: -1,
	}
}

func (scr *screen_prodchoose) View() string {
	s := "What would you like to buy?\n\n"

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

func (scr *screen_prodchoose) Update(msg tea.Msg) tea.Cmd {
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
				mainmodel.MainModel.LoadScreen(ProdConfirm(scr.choices[scr.selected]))
				// mainmodel.MainModel.LoadModel(prodconfirm.CreateModel(m.choices[m.selected]))
			}
		}
	}

	return nil
}
