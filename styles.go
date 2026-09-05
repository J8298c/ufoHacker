package main

import "charm.land/lipgloss/v2"

var background = lipgloss.Color("#0A0A0A")

var scrambledStyles = lipgloss.NewStyle().Foreground(lipgloss.Color("#8A5A00"))

var decodedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB000"))

var panelStyle = lipgloss.NewStyle().
	Border(lipgloss.DoubleBorder()).
	BorderForeground(lipgloss.Color("#8A5A00")).
	Background(background).
	Padding(1, 2)

var screenStyle = lipgloss.NewStyle().
	Background(background).
	Padding(1, 2)

var selectedItemStyle = lipgloss.NewStyle().
	Foreground(background).
	Background(lipgloss.Color("#FFB000"))

var bannerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#8A5A00"))

const asciiBanner = `
              _____
         .-'"     "'-.
      .-'             '-.
    .'                   '.
   /                       \
  |    O     O     O     O  |
   '-.___________________.-'
          |    |    |
          '----+----'

  N . U . F . O . R . C.   A R C H I V E
`
