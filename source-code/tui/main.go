package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	command  string
	args     string
}

func initialModel() model {
	return model{
		choices:  []string{"install", "remove", "update", "upgrade", "search"},
		selected: make(map[int]struct{}),
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
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.command == "" {
				m.command = m.choices[m.cursor]
				m.args = "" // Prompt for args later if needed
				return m, runCommand(m.command, m.args)
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Select APT command:\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress q to quit.\n"
	return lipgloss.NewStyle().Margin(1, 2).Render(s)
}

func runCommand(command, args string) tea.Cmd {
	return func() tea.Msg {
		var aptCmd string
		var aptArgs []string

		switch command {
		case "install":
			aptCmd = "apt"
			aptArgs = append([]string{"install", "-y"}, strings.Split(args, " ")...)
		case "remove":
			aptCmd = "apt"
			aptArgs = append([]string{"remove", "-y"}, strings.Split(args, " ")...)
		case "update":
			aptCmd = "apt"
			aptArgs = []string{"update"}
		case "upgrade":
			aptCmd = "apt"
			aptArgs = []string{"upgrade", "-y"}
		case "search":
			aptCmd = "apt"
			aptArgs = append([]string{"search"}, strings.Split(args, " ")...)
		}

		cmd := exec.Command("sudo", append([]string{aptCmd}, aptArgs...)...)
		output, _ := cmd.CombinedOutput()
		fmt.Println(string(output))
		return nil
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
