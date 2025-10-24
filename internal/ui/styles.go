// file: internal/ui/styles.go
package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Deus Ex color palette - orange/amber cyberpunk theme
	primaryOrange = lipgloss.Color("#FF9800")
	accentOrange  = lipgloss.Color("#FFB74D")
	darkOrange    = lipgloss.Color("#F57C00")
	cyberGold     = lipgloss.Color("#FFA726")
	warningAmber  = lipgloss.Color("#FF6F00")

	textLight = lipgloss.Color("#E0E0E0")
	textMuted = lipgloss.Color("#9E9E9E")
	textDark  = lipgloss.Color("#616161")

	bgDark   = lipgloss.Color("#0A0A0A")
	bgMedium = lipgloss.Color("#1A1A1A")
	bgLight  = lipgloss.Color("#2A2A2A")

	successGreen = lipgloss.Color("#66BB6A")
	errorRed     = lipgloss.Color("#EF5350")

	baseStyle = lipgloss.NewStyle().
		Foreground(textLight).
		Background(bgDark)

	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(primaryOrange).
		Background(bgMedium).
		Padding(0, 2).
		MarginBottom(1).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(darkOrange)

	subtitleStyle = lipgloss.NewStyle().
		Foreground(accentOrange).
		Italic(true).
		MarginBottom(1)

	inputStyle = lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(primaryOrange).
		Padding(0, 1).
		Foreground(textLight).
		Background(bgMedium)

	inputPromptStyle = lipgloss.NewStyle().
		Foreground(cyberGold).
		Bold(true)

	listStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(darkOrange).
		Padding(1, 2).
		MarginTop(1).
		Background(bgMedium)

	selectedItemStyle = lipgloss.NewStyle().
		Foreground(bgDark).
		Background(primaryOrange).
		Bold(true).
		Padding(0, 1)

	itemStyle = lipgloss.NewStyle().
		Foreground(textLight).
		Padding(0, 1)

	itemNumberStyle = lipgloss.NewStyle().
		Foreground(cyberGold).
		Bold(true).
		Width(4)

	episodeTitleStyle = lipgloss.NewStyle().
		Foreground(accentOrange).
		Bold(true)

	episodeMetaStyle = lipgloss.NewStyle().
		Foreground(textMuted).
		Italic(true)

	episodeDescStyle = lipgloss.NewStyle().
		Foreground(textLight).
		MarginTop(1).
		Width(80)

	statusStyle = lipgloss.NewStyle().
		Foreground(textMuted).
		MarginTop(1).
		MarginBottom(1).
		Italic(true).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(textDark)

	errorStyle = lipgloss.NewStyle().
		Foreground(errorRed).
		Bold(true).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(errorRed)

	successStyle = lipgloss.NewStyle().
		Foreground(successGreen).
		Bold(true)

	loadingStyle = lipgloss.NewStyle().
		Foreground(accentOrange).
		Bold(true).
		Padding(1, 2)

	helpStyle = lipgloss.NewStyle().
		Foreground(textMuted).
		MarginTop(1).
		Padding(1, 0).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(textDark)

	helpKeyStyle = lipgloss.NewStyle().
		Foreground(primaryOrange).
		Bold(true)

	containerStyle = lipgloss.NewStyle().
		Padding(2, 3).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(primaryOrange).
		Background(bgDark)

	dividerStyle = lipgloss.NewStyle().
		Foreground(darkOrange).
		MarginTop(0).
		MarginBottom(0)

	badgeStyle = lipgloss.NewStyle().
		Foreground(bgDark).
		Background(cyberGold).
		Bold(true).
		Padding(0, 1).
		MarginRight(1)

	metaStyle = lipgloss.NewStyle().
		Foreground(textMuted).
		Italic(true)
)
