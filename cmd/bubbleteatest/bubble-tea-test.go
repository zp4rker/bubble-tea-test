package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zp4rker/bubbleteatest/internal/bubbleteatest/mainmodel"
	"github.com/zp4rker/bubbleteatest/internal/bubbleteatest/screens"
)

func main() {
	mainmodel.MainModel.LoadScreen(screens.ProdChoose())

	p := tea.NewProgram(mainmodel.MainModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Encountered an error: %v", err)
		os.Exit(1)
	}
}
