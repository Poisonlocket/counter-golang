package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	err := tea.NewProgram(&Model{}, tea.WithAltScreen()).Start()
	if err != nil {
		return
	}
}

type Model struct {
	count int
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
			os.Exit(1)

		case "up":
			m.count++

		case "down":
			m.count--
		}

	}
	return m, nil
}

func (m *Model) View() string {
	return fmt.Sprintf("Count: %d\n\n↑ Increase,\t↓ Decrease", m.count)
}
