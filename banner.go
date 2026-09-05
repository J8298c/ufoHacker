package main

import (
	"strings"

	"charm.land/lipgloss/v2"
	figure "github.com/common-nighthawk/go-figure"
)

func figlet(text, font string) string {
	return strings.Join(figure.NewFigure(text, font, true).Slicify(), "\n")
}

func buildBanner(width int) string {
	title := titleStyle.Render(figlet("NUFORC", "smslant"))
	block := blockTitleStyle.Render(figlet("CLASSIFIED", "doom"))
	tagline := taglineStyle.Render(">> NATIONAL UFO REPORTING CENTER — RESTRICTED ACCESS")
	warning := warningBoxStyle.Render(
		"WARNING: UNAUTHORIZED ACCESS TO THIS TERMINAL IS LOGGED\n" +
			"ALL TRANSMISSIONS ARE ENCRYPTED — DECRYPTION KEY REQUIRED",
	)

	banner := lipgloss.JoinVertical(lipgloss.Center, title, block, tagline, warning)
	return lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(banner)
}
