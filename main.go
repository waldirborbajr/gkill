package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

type fileItem string

func (f fileItem) Title() string { return string(f) }
func (f fileItem) Description() string {
	return ""
}
func (f fileItem) FilterValue() string { return string(f) }

func initialModel() model {
	files, err := getFilesInCurrentDir()
	if err != nil {
		log.Fatal(err)
	}

	items := []list.Item{}
	for _, file := range files {
		items = append(items, fileItem(file))
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select a file"

	return model{
		list: l,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(fileItem)
			if ok {
				m.choice = string(i)
				createSymLink(string(i))
			}
			return m, tea.Quit

		default:
			var cmd tea.Cmd
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		}

	default:
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.quitting {
		return "\n"
	}
	str := "\n" + m.list.View() + "\n\n"

	if m.choice != "" {
		str += fmt.Sprintf(
			"A symbolic link for %s has been created in your home directory.\n",
			m.choice,
		)
	}

	str += "\nPress q to quit.\n"

	return str
}

func getFilesInCurrentDir() ([]string, error) {
	files, err := os.ReadDir("./")
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}

	return filenames, nil
}

func createSymLink(filename string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	oldPath, err := filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
	}

	newPath := filepath.Join(homeDir, filename)

	err = os.Symlink(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := tea.NewProgram(initialModel())
	app.Run()
}
