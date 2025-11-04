package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	msg string
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit
        }
    }

    return m, nil
}

func (m model) View() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingLeft(2).
		PaddingRight(2)

	sectionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#D6D6D6"))

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#999999")).
		MarginTop(1)

	welcome := titleStyle.Render("Welcome to TermNote! >_")
	view := sectionStyle.Render("View coming soon...")
	help := helpStyle.Render("Ctrl+N: new file . Ctrl+L: list . Esc: back/save . Ctrl+S: save . Ctrl+Q: quit")

	return fmt.Sprintf("%s\n\n%s\n\n%s", welcome, view, help)
}

func initializeModel() model {
	return model{
		msg: "TermNote - A terminal note app!",
	}
}
func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}