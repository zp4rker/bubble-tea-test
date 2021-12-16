package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/mainmodel"
	"github.com/zp4rker/bubble-tea-test/internal/bubble-tea-test/screens"
)

func main() {
	mainmodel.MainModel.LoadScreen(screens.ProdChoose())

	p := tea.NewProgram(mainmodel.MainModel, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Encountered an error: %v", err)
		os.Exit(1)
	}
}
