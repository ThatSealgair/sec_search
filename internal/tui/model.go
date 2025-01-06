package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ThatSealgair/sec_search/internal/converter"
	"github.com/ThatSealgair/sec_search/internal/platform"
)

type Model struct {
	list          list.Model
	textInput     textinput.Model
	state         state
	selectedFrom  platform.Platform
	selectedTo    platform.Platform
	result        string
	err           error
	converter     *converter.QueryConverter
	width, height int
}

type item struct {
	title    string
	platform platform.Platform
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.title }

func NewModel() Model {
	// Initialize text input
	ti := textinput.New()
	ti.Placeholder = "Enter your search query..."
	ti.CharLimit = 512
	ti.Width = 80
	ti.Prompt = "│ "
	ti.PromptStyle = inputPromptStyle
	ti.TextStyle = inputTextStyle

	// Initialize platform list
	items := []list.Item{
		item{title: "Shodan", platform: platform.Shodan},
		item{title: "Censys", platform: platform.Censys},
		item{title: "Hunter", platform: platform.Hunter},
	}

	// Custom list delegate
	delegate := list.NewDefaultDelegate()
	delegate.Styles.NormalTitle = listItemStyle
	delegate.Styles.SelectedTitle = selectedItemStyle
	delegate.Styles.DimmedTitle = listItemStyle.Copy().Foreground(lipgloss.Color(dragonBlack4))

	l := list.New(items, delegate, 0, 0)
	l.Title = "Select Platform"
	l.Styles.Title = listHeaderStyle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return Model{
		list:      l,
		textInput: ti,
		state:     stateSelectFromPlatform,
		converter: converter.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		if m.width > 0 {
			m.textInput.Width = m.width - 4
		}

		headerHeight := 4
		footerHeight := 2
		verticalMarginHeight := headerHeight + footerHeight
		m.list.SetSize(m.width-4, m.height-verticalMarginHeight)

		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			switch m.state {
			case stateSelectFromPlatform:
				if i, ok := m.list.SelectedItem().(item); ok {
					m.selectedFrom = i.platform
					m.state = stateSelectToPlatform
					m.list.Title = "Select Target Platform"
				}

			case stateSelectToPlatform:
				if i, ok := m.list.SelectedItem().(item); ok {
					m.selectedTo = i.platform
					m.state = stateEnterQuery
					m.textInput.Focus()
					return m, textinput.Blink
				}

			case stateEnterQuery:
				query := strings.TrimSpace(m.textInput.Value())
				if query != "" {
					result, err := m.converter.Convert(query, m.selectedFrom, m.selectedTo)
					if err != nil {
						m.err = err
					} else {
						m.result = result
						m.textInput.Blur()
					}
					m.state = stateShowResult
				}

			case stateShowResult:
				// Reset for new conversion
				m.state = stateSelectFromPlatform
				m.selectedFrom = ""
				m.selectedTo = ""
				m.result = ""
				m.err = nil
				m.textInput.Reset()
				m.list.Title = "Select Platform"
				m.list.ResetSelected()
			}
		}
	}

	// Handle component updates based on state
	switch m.state {
	case stateSelectFromPlatform, stateSelectToPlatform:
		var listCmd tea.Cmd
		m.list, listCmd = m.list.Update(msg)
		return m, listCmd

	case stateEnterQuery:
		var inputCmd tea.Cmd
		m.textInput, inputCmd = m.textInput.Update(msg)
		return m, inputCmd
	}

	return m, cmd
}

func (m Model) View() string {
	var s strings.Builder

	// Header
	s.WriteString(titleStyle.Render("Security Search Query Converter"))
	s.WriteString("\n\n")

	// Content based on state
	switch m.state {
	case stateSelectFromPlatform:
		s.WriteString(subtitleStyle.Render("Select source platform:"))
		s.WriteString("\n\n")
		s.WriteString(MainLayout(m.list.View()))

	case stateSelectToPlatform:
		s.WriteString(subtitleStyle.Render(fmt.Sprintf(
			"Converting from %s",
			highlightStyle.Render(string(m.selectedFrom)),
		)))
		s.WriteString("\n\n")
		s.WriteString(MainLayout(m.list.View()))

	case stateEnterQuery:
		s.WriteString(subtitleStyle.Render(fmt.Sprintf(
			"Converting: %s → %s",
			highlightStyle.Render(string(m.selectedFrom)),
			highlightStyle.Render(string(m.selectedTo)),
		)))
		s.WriteString("\n\n")
		s.WriteString(infoStyle.Render("Enter your query:"))
		s.WriteString("\n")
		s.WriteString(m.textInput.View())

	case stateShowResult:
		s.WriteString(subtitleStyle.Render(fmt.Sprintf(
			"Converting: %s → %s",
			highlightStyle.Render(string(m.selectedFrom)),
			highlightStyle.Render(string(m.selectedTo)),
		)))
		s.WriteString("\n\n")

		if m.err != nil {
			s.WriteString(errorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
		} else {
			s.WriteString(infoStyle.Render("Original query:"))
			s.WriteString("\n")
			s.WriteString(resultStyle.Render(m.textInput.Value()))
			s.WriteString("\n\n")
			s.WriteString(infoStyle.Render("Converted query:"))
			s.WriteString("\n")
			s.WriteString(resultStyle.Render(m.result))
		}

		s.WriteString("\n\n")
		s.WriteString(statusBarStyle.Render("Press Enter for new conversion • Esc to quit"))
	}

	// Footer
	s.WriteString("\n")
	s.WriteString(statusBarStyle.Render(fmt.Sprintf(
		"%s • %s",
		time.Now().Format("15:04:05"),
		getStatusText(m.state),
	)))

	return s.String()
}

func getStatusText(s state) string {
	switch s {
	case stateSelectFromPlatform:
		return "Select source platform with arrow keys"
	case stateSelectToPlatform:
		return "Select target platform with arrow keys"
	case stateEnterQuery:
		return "Enter your query and press Enter"
	case stateShowResult:
		return "Press Enter for new conversion"
	default:
		return ""
	}
}
