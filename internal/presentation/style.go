package presentation

import "github.com/charmbracelet/lipgloss"

var (
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	footerStyle = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("241"))
	logoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#800080"))
)
