package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mbarlow/bubble-skeleton/internal/config"
	"github.com/mbarlow/bubble-skeleton/internal/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cfg := config.New()
	
	if cfg.Debug {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	m := models.New(cfg)
	p := tea.NewProgram(m, tea.WithAltScreen())
	
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}