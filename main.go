package main

//Step1
import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"log"
	"os"
)

// Step2
type model struct {
	newFileTextInput       textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	noteTextArea           textarea.Model
}

// Constants
var (
	foregroundColor = lipgloss.Color("16")
	backgroundColor = lipgloss.Color("99")
)

// Component styling
var (
	cursorStyle = lipgloss.NewStyle().Foreground(backgroundColor)
)

// Paths
var (
	vaultDir string
)

// Golang init function that is called just before main() fn
func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting user home directory", err)
	}

	//fmt.Printf("Golang Init: ", homeDir)

	//Files will be created here
	vaultDir = fmt.Sprintf("%s\\.totion", homeDir)

	//fmt.Printf("Golang Init: ", vaultDir)

}

// Step3
func initializeModel() model {

	fmt.Printf("Step3", vaultDir)
	//Create directory if not exist
	err := os.MkdirAll(vaultDir, 0750)
	if err != nil {
		log.Fatal(err)
	}

	//Initialize Model

	//Text Input
	t1 := textinput.New()
	t1.Placeholder = "What would you call it?"
	t1.Focus()
	t1.CharLimit = 156
	t1.Cursor.Style = cursorStyle

	//Text Area
	t2 := textarea.New()
	t2.Placeholder = "// :Todo"
	t2.Focus()
	t2.Cursor.Style = cursorStyle

	return model{
		newFileTextInput:       t1,
		createFileInputVisible: false, //by default, it is false when we call Ctrl+N then it's True
		noteTextArea:           t2,
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
		case "enter":

			//create file
			filename := m.newFileTextInput.Value()
			//fmt.Printf("Update method - Filename: ", filename)
			if filename != "" {
				filePath := fmt.Sprintf("%s\\%s.md", vaultDir, filename)
				//fmt.Printf("Update method - Filepath: ", filePath)

				//return the statistics of file or error
				_, err := os.Stat(filePath) // if this file not found it will throw error
				//if error found means file path not exist
				if err == nil {
					return m, nil
				}

				//fmt.Printf("Update method - FileInfo: ", fileInfo)
				//create the file with existing folder
				file, err := os.Create(filePath)
				if err != nil {
					log.Fatalf("%v", err)
				}

				//fmt.Printf("Update method - File: ", file)

				//hide file input field
				m.createFileInputVisible = false

				//clear the text from file input field
				m.newFileTextInput.SetValue("")

				//update state so that we can show text editor are
				m.currentFile = file

				return m, nil
			}
			//text editor -> Write in the file descriptor and close it.
		case "ctrl+s":

			if m.currentFile == nil { //if no file editor is opened
				break
			}
			if err := m.currentFile.Truncate(0); err != nil {
				fmt.Println("can not save the file :(", err)
				return m, nil
			}

			//offset is line number of height and whence is width hence means which point you are.
			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("can not save the file :(", err)
				return m, nil
			}
			if _, err := m.currentFile.WriteString(m.noteTextArea.Value()); err != nil {
				//fmt.Println("Saving the file :)")
				return m, nil
			}

			if err := m.currentFile.Close(); err != nil {
				fmt.Println("can not close the file :(", err)
			}

			//update the state
			m.currentFile = nil
			m.noteTextArea.SetValue("")

			return m, nil
		}

	}

	//Update the state of file input field
	if m.createFileInputVisible {
		//when file input is visible, update user message to our model file input
		m.newFileTextInput, cmd = m.newFileTextInput.Update(msg)
	}
	//Update the state of text editor
	if m.currentFile != nil {
		m.noteTextArea, cmd = m.noteTextArea.Update(msg)
	}
	return m, cmd
}

// Step 6
func (m model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(foregroundColor).
		Background(backgroundColor).
		PaddingLeft(2).PaddingRight(4)

	welcomeMessage := style.Render("Welcome to Totion 📝")
	help := "Ctrl+N: new file • Ctrl+L: list • Esc: back/save • Ctrl+S: save • Ctrl+Q: quit"
	view := ""

	if m.createFileInputVisible {
		view = m.newFileTextInput.View() //Call Text Input's View Method to Show Input Field
	}

	//if there is file created
	if m.currentFile != nil {
		view = m.noteTextArea.View()

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
