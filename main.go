package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/j8298c/ufoHacker/records"
)

const ShiftKey = 13
const maxAttempts = 3

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
	guessInput     string
	attempts       int
	solved         bool
	feedback       string
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
			switch {
			case m.screen == "list":
				m.selectedRecord = m.records[m.cursor]
				m.screen = "detail"
				m.guessInput = ""
				m.attempts = maxAttempts
				m.solved = false
				m.feedback = ""
			case m.screen == "detail" && !m.solved:
				guess, err := strconv.Atoi(m.guessInput)
				if err == nil && guess == ShiftKey {
					m.solved = true
					m.feedback = "ACCESS GRANTED — TRANSMISSION DECRYPTED"
				} else {
					m.attempts--
					if m.attempts <= 0 {
						m.solved = true
						m.feedback = "ATTEMPTS EXHAUSTED — AUTO-DECRYPTING TRANSMISSION"
					} else {
						m.feedback = fmt.Sprintf("INCORRECT KEY — ATTEMPTS REMAINING: %d", m.attempts)
					}
				}
				m.guessInput = ""
			}
		case "backspace":
			if m.screen == "detail" && !m.solved && len(m.guessInput) > 0 {
				m.guessInput = m.guessInput[:len(m.guessInput)-1]
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
		default:
			if m.screen == "detail" && !m.solved && len(m.guessInput) < 2 &&
				len(msg.Text) == 1 && msg.Text[0] >= '0' && msg.Text[0] <= '9' {
				m.guessInput += msg.Text
			}
		}
	}
	return m, nil
}

// record loader find records upon state choice'

// setup view
func (m model) View() tea.View {
	var content string
	if m.screen == "detail" {
		content = m.renderDetail()
	} else {
		content = m.renderList()
	}
	banner := buildBanner(m.width)
	body := lipgloss.JoinVertical(lipgloss.Left, banner, content)
	screen := screenStyle.Width(m.width).Height(m.height).Render(body)
	v := tea.NewView(screen)
	v.AltScreen = true
	return v
}

func (m model) renderDetail() string {
	if m.solved {
		r := m.selectedRecord
		decoded := CaesarShift(r.Summary, -ShiftKey)
		return strings.Join([]string{
			decodedStyle.Render(m.feedback),
			"",
			decodedStyle.Render(fmt.Sprintf("LOCATION: %s, %s", r.City, r.State)),
			decodedStyle.Render(fmt.Sprintf("COORDINATES: %.4f, %.4f", r.CityLatitude, r.CityLongitude)),
			decodedStyle.Render(fmt.Sprintf("SHAPE: %s", r.Shape)),
			"",
			decodedStyle.Render(decoded),
		}, "\n")
	}

	prompt := fmt.Sprintf("ENTER 2-DIGIT SHIFT KEY: %s_", m.guessInput)
	lines := []string{
		scrambledStyles.Render(m.selectedRecord.Summary),
		"",
		decodedStyle.Render(prompt),
	}
	if m.feedback != "" {
		lines = append(lines, decodedStyle.Render(m.feedback))
	}
	return strings.Join(lines, "\n")
}

func (m model) renderList() string {
	var lines []string
	for i, r := range m.records {
		label := fmt.Sprintf("%s, %s [%.4f, %.4f] (%s)", r.City, r.State, r.CityLatitude, r.CityLongitude, r.Shape)
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
