package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/j8298c/ufoHacker/records"
)

// model for applications
type model struct {
	states         []string
	stateChoice    string
	records        []records.Record
	selectedRecord records.Record
	width          int
	height         int
	screen         string
	cursor         int
}

// record structs

func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// create initial model
func initializeModel() model {
	return model{
		states: []string{
			"NY",
		},
		records: records.Records,
		screen:  "list",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// update command
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.screen == "list" && m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.screen == "list" && m.cursor < len(m.records)-1 {
				m.cursor++
			}
		case "enter":
			if m.screen == "list" {
				m.selectedRecord = m.records[m.cursor]
				m.screen = "detail"
			}
		case "esc":
			if m.screen == "detail" {
				m.screen = "list"
			}
		case "q":
			if m.screen == "list" {
				return m, tea.Quit
			}
			m.screen = "list"
		}
	}
	return m, nil
}

// record loader find records upon state choice'

// setup view
func (m model) View() tea.View {
	var content string
	if m.screen == "detail" {
		content = scrambledStyles.Render(m.selectedRecord.Summary)
	} else {
		content = m.renderList()
	}
	banner := bannerStyle.Width(m.width).Align(lipgloss.Center).Render(asciiBanner)
	body := lipgloss.JoinVertical(lipgloss.Left, banner, content)
	screen := screenStyle.Width(m.width).Height(m.height).Render(body)
	v := tea.NewView(screen)
	v.AltScreen = true
	return v
}

func (m model) renderList() string {
	var lines []string
	for i, r := range m.records {
		label := fmt.Sprintf("%s, %s (%s)", r.City, r.State, r.Shape)
		if i == m.cursor {
			lines = append(lines, selectedItemStyle.Render("> "+label))
		} else {
			lines = append(lines, decodedStyle.Render("  "+label))
		}
	}
	return strings.Join(lines, "\n")
}

// decode message
func CaesarShift(s string, n int) string {
	shift := byte(((n % 26) + 26) % 26)
	out := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= 'a' && c <= 'z':
			out[i] = 'a' + (c-'a'+shift)%26
		case c >= 'A' && c <= 'Z':
			out[i] = 'A' + (c-'A'+shift)%26
		default:
			out[i] = c
		}
	}
	return string(out)
}
