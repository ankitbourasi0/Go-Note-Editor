package main

//Step1
import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

// Step2
type model struct {
	newFileTextInput       textinput.Model
	createFileInputVisible bool
}

// Step3
func initializeModel() model {
	//Initialize Model
	ti := textinput.New()
	ti.Placeholder = "What would you like to do?"
	ti.Focus()
	ti.CharLimit = 156

	return model{
		newFileTextInput:       ti,
		createFileInputVisible: false, //by default, it is false when we call Ctrl+N then it's True
	}

}

// Step4
func (m model) Init() tea.Cmd {
	return nil
}

// Step5
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	/*Cmd is an IO operation that returns a message when it's complete. If it's nil it's considered a no-op.
	Use it for things like HTTP requests, timers, saving and loading from disk, and so on.*/
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		//New file
		case "ctrl+n":
			m.createFileInputVisible = true

			return m, nil
		}
	}

	if m.createFileInputVisible {
		//when file input is visible, update user message to our model file input
		m.newFileTextInput, cmd = m.newFileTextInput.Update(msg)
	}
	return m, cmd
}

// Step6
func (m model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("99")).
		PaddingLeft(2).PaddingRight(4)

	welcomeMessage := style.Render("Welcome to Totion📝")
	help := "Ctrl+N: new file • Ctrl+L: list • Esc: back/save • Ctrl+S: save • Ctrl+Q: quit"
	view := ""

	if m.createFileInputVisible {
		view = m.newFileTextInput.View() //Call Text Input's View Method to Show Input Field
	}
	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcomeMessage, view, help)
}

// Step7
func main() {
	program := tea.NewProgram(initializeModel())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Ahh! there's been an error: %v", err)
		os.Exit(1)
	}
}
