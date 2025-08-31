package views

import (
	"fmt"
	"strings"

	"github.com/mbarlow/bubble-skeleton/internal/styles"
	"github.com/charmbracelet/lipgloss"
)

type View struct {
	width  int
	height int
	styles *styles.Styles
}

func New() *View {
	return &View{
		styles: styles.New(),
	}
}

func (v *View) SetSize(width, height int) {
	v.width = width
	v.height = height
}

func (v *View) Loading() string {
	content := v.styles.Title.Render("Loading...")
	return lipgloss.Place(
		v.width,
		v.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (v *View) Main() string {
	var sections []string

	header := v.renderHeader()
	sections = append(sections, header)

	content := v.renderContent()
	sections = append(sections, content)

	footer := v.renderFooter()
	sections = append(sections, footer)

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (v *View) Help() string {
	helpText := []string{
		v.styles.Title.Render("Help"),
		"",
		v.styles.Help.Render("Keyboard Shortcuts:"),
		"",
		fmt.Sprintf("  %s - Show this help", v.styles.Key.Render("h, ?")),
		fmt.Sprintf("  %s - Enter input mode", v.styles.Key.Render("i")),
		fmt.Sprintf("  %s - Refresh", v.styles.Key.Render("r")),
		fmt.Sprintf("  %s - Quit", v.styles.Key.Render("q, ctrl+c")),
		fmt.Sprintf("  %s - Return to main view", v.styles.Key.Render("esc")),
		"",
		v.styles.Subtle.Render("Press any key to return"),
	}

	content := lipgloss.JoinVertical(lipgloss.Left, helpText...)
	
	return lipgloss.Place(
		v.width,
		v.height,
		lipgloss.Center,
		lipgloss.Center,
		v.styles.Box.Render(content),
	)
}

func (v *View) Input(value string, cursor int) string {
	var sections []string

	header := v.styles.Title.Render("Input Mode")
	sections = append(sections, header)
	sections = append(sections, "")

	prompt := v.styles.Prompt.Render("Enter value: ")
	
	displayValue := value
	if cursor < len(value) {
		before := value[:cursor]
		at := string(value[cursor])
		after := ""
		if cursor+1 < len(value) {
			after = value[cursor+1:]
		}
		displayValue = before + v.styles.Cursor.Render(at) + after
	} else {
		displayValue = value + v.styles.Cursor.Render(" ")
	}
	
	inputLine := prompt + displayValue
	sections = append(sections, inputLine)
	sections = append(sections, "")
	
	help := v.styles.Subtle.Render("Press Enter to submit, Esc to cancel")
	sections = append(sections, help)

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)
	
	return lipgloss.Place(
		v.width,
		v.height,
		lipgloss.Center,
		lipgloss.Center,
		v.styles.Box.Render(content),
	)
}

func (v *View) renderHeader() string {
	title := v.styles.Title.Render("Bubble Tea Skeleton")
	subtitle := v.styles.Subtitle.Render("A reusable template for Bubble Tea applications")
	
	header := lipgloss.JoinVertical(lipgloss.Center, title, subtitle)
	
	return v.styles.Header.
		Width(v.width).
		Render(header)
}

func (v *View) renderContent() string {
	content := []string{
		v.styles.Section.Render("Welcome!"),
		"",
		"This is a skeleton Bubble Tea application that you can use as a starting point for your projects.",
		"",
		"Features:",
		"  • Modular structure with separate packages",
		"  • Styled components using Lipgloss",
		"  • Command pattern for actions",
		"  • Configuration management",
		"  • Input handling",
		"  • Help system",
		"",
		v.styles.Subtle.Render("Press 'h' for help"),
	}
	
	joined := lipgloss.JoinVertical(lipgloss.Left, content...)
	
	contentHeight := v.height - 8
	if contentHeight < 1 {
		contentHeight = 1
	}
	
	return v.styles.Content.
		Width(v.width - 4).
		Height(contentHeight).
		Render(joined)
}

func (v *View) renderFooter() string {
	left := v.styles.StatusBarLeft.Render("Ready")
	right := v.styles.StatusBarRight.Render(fmt.Sprintf("%dx%d", v.width, v.height))
	
	gap := v.width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 0 {
		gap = 0
	}
	
	footer := left + strings.Repeat(" ", gap) + right
	
	return v.styles.StatusBar.
		Width(v.width).
		Render(footer)
}