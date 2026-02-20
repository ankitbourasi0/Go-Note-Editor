package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	msg string
}

func (m model) Init() tea.Cmd {
	return nil
}

// Update method usually receive a message type, but we can use type assertion also
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	//Is it a key press?
	case tea.KeyMsg:

		//Great, what was the actual press?
		switch msg.String() {
		//Exit the application when these keys pressed.
		case "ctrl+c", "q":
			fmt.Println("\n\nuser pressed", msg.String(), "Application Exit")
			return m, tea.Quit

		}
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
func (m model) View() string {
	return m.msg
}

func initializeMode() model {
	return model{
		msg: "Initializing...",
	}
}

func main() {
	p := tea.NewProgram(initializeMode())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ahh! there's been an error: %v", err)
		os.Exit(1)
	}
}
