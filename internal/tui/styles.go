// internal/tui/styles.go
package tui

import "github.com/charmbracelet/lipgloss"

// Kanagawa Dragon theme colors
const (
	// Base colors
	dragonBlack0 = "#0D0C0C" // Darker background
	dragonBlack1 = "#12120F" // Dark background
	dragonBlack2 = "#1D1C19" // Selection background
	dragonBlack3 = "#363646" // Dark foreground
	dragonBlack4 = "#54546D" // Comments
	dragonBlack5 = "#625E5A" // Dark steel
	dragonBlack6 = "#72696A" // Gray steel

	// Special colors
	dragonRed    = "#C4746E" // Error and warning
	dragonOrange = "#B6927B" // Numbers and constants
	dragonYellow = "#DCA561" // Functions and titles
	dragonGreen  = "#8A9A7B" // Strings and success
	dragonBlue   = "#8BA4B0" // Keywords and info
	dragonViolet = "#957FB8" // Special keywords
	dragonAqua   = "#7AA89F" // Operators and hints
	dragonPink   = "#D27E99" // Parameters and variables

	// UI elements
	borderRadius = 1
	padding      = 1
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(dragonYellow)).
			Background(lipgloss.Color(dragonBlack1)).
			Padding(padding).
			MarginLeft(2)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dragonBlue)).
			Background(lipgloss.Color(dragonBlack1)).
			MarginLeft(2)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dragonRed)).
			Background(lipgloss.Color(dragonBlack1)).
			MarginLeft(2)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dragonBlack4)).
			Background(lipgloss.Color(dragonBlack1)).
			MarginLeft(2)

	highlightStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dragonAqua)).
			Background(lipgloss.Color(dragonBlack1))

	resultStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(dragonViolet)).
			Foreground(lipgloss.Color(dragonGreen)).
			Background(lipgloss.Color(dragonBlack2)).
			Padding(1, 2).
			MarginLeft(2)

	// List styles
	listHeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dragonPink)).
			Background(lipgloss.Color(dragonBlack1)).
			Padding(0, 1)

	listItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dragonBlack6)).
			Background(lipgloss.Color(dragonBlack1)).
			Padding(0, 1)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(dragonYellow)).
				Background(lipgloss.Color(dragonBlack2)).
				Bold(true).
				Padding(0, 1)

	// Input styles
	inputPromptStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(dragonOrange)).
				Background(lipgloss.Color(dragonBlack1))

	inputTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dragonAqua)).
			Background(lipgloss.Color(dragonBlack2)).
			Padding(0, 1)

	// Status bar style
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dragonBlack4)).
			Background(lipgloss.Color(dragonBlack1)).
			Padding(0, 1)
)

// Helper function for consistent main layout
func MainLayout(content string) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color(dragonBlack1)).
		Padding(1, 2).
		Render(content)
}
