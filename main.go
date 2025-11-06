package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
	vaultDir    string
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error gettting home directory", err)
	}
	vaultDir = fmt.Sprintf("%s/.termnote", homeDir)
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	newFileInput           textinput.Model
	noteTextArea           textarea.Model
	createFileInputVisible bool
	currentFile            *os.File

	list        list.Model
	showingList bool
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v-5)

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.createFileInputVisible {
				m.createFileInputVisible = false
			}

			if m.currentFile != nil {
				m.noteTextArea.SetValue("")
				m.currentFile = nil
			}

			if m.showingList {
				if m.list.FilterState() == list.Filtering {
					break
				}
				m.showingList = false
			}
			return m, nil
		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil
		case "ctrl+l":
			noteList := listFiles()
			m.list.SetItems(noteList)
			m.showingList = true
			return m, nil
		case "ctrl+s":
			if m.currentFile == nil {
				break
			}

			if err := m.currentFile.Truncate(0); err != nil {
				fmt.Println("Cannot save the file :(")
				return m, nil
			}

			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("Cannot save the file :(")
				return m, nil
			}

			// write the content of the text area to the file
			if _, err := m.currentFile.WriteString(m.noteTextArea.Value()); err != nil {
				fmt.Println("Cannot write to the file :(")
				return m, nil
			}

			if err := m.currentFile.Close(); err != nil {
				fmt.Println("Cannot close the file :(")
			}

			m.currentFile = nil
			m.noteTextArea.SetValue("")

			return m, nil
		case "enter":
			if m.currentFile != nil {
				break
			}

			if m.showingList {
				item, ok := m.list.SelectedItem().(item)
				if ok {
					filePath := fmt.Sprintf("%s/%s", vaultDir, item.title)

					content, err := os.ReadFile(filePath)
					if err != nil {
						log.Printf("Error reading file: %v", err)
						return m, nil
					}

					m.noteTextArea.SetValue(string(content))
					f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
					if err != nil {
						log.Printf("Error opening file: %v", err)
						return m, nil
					}

					m.currentFile = f
					m.showingList = false
				}

				return m, nil
			}

			fileName := m.newFileInput.Value()
			if fileName != "" {
				filePath := fmt.Sprintf("%s/%s.md", vaultDir, fileName)

				if _, err := os.Stat(filePath); err == nil {
					return m, nil
				}

				f, err := os.Create(filePath)
				if err != nil {
					log.Fatal("Error creating file", err)
				}

				m.currentFile = f
				m.createFileInputVisible = false
				m.newFileInput.SetValue("")
			}
			return m, nil
		}
	}

	if m.createFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}

	if m.currentFile != nil {
		m.noteTextArea, cmd = m.noteTextArea.Update(msg)
	}

	if m.showingList {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingLeft(2).
		PaddingRight(2)

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#999999"))

	welcome := titleStyle.Render("Welcome to TermNote! >_")
	help := helpStyle.Render("Ctrl+n: new file . Ctrl+l: list . Esc: back . Ctrl+s: save . Ctrl+q: quit")

	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.noteTextArea.View()
	}

	if m.showingList {
		view = m.list.View()
	}

	return fmt.Sprintf("%s\n\n%s\n\n%s", welcome, view, help)
}

func initializeModel() model {
	err := os.MkdirAll(vaultDir, 0750)
	if err != nil {
		log.Fatal("Error creating vault directory", err)
	}

	// initialize new file input
	ti := textinput.New()
	ti.Placeholder = "Enter the name of your new note..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = cursorStyle

	// initialize note text area
	ta := textarea.New()
	ta.Placeholder = "Start writing your note..."
	ta.Focus()

	// list
	noteList := listFiles()
	finalList := list.New(noteList, list.NewDefaultDelegate(), 0, 0)
	finalList.Title = "TermNote - Your Notes"
	finalList.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7D56F4"))

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		noteTextArea:           ta,
		list: 				 finalList,
	}
}
func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func listFiles() []list.Item {
	items := make([]list.Item, 0)

	entries, err := os.ReadDir(vaultDir)
	if err != nil {
		log.Fatal("Error reading notes list")
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			modifiedTime := info.ModTime().Format("2006-01-02 15:04:05")
			items = append(items, item{
				title: entry.Name(),
				desc:  fmt.Sprintf("Last Modified: %s", modifiedTime),
			})
		}
	}

	return items
}
