package styles

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Title          lipgloss.Style
	Subtitle       lipgloss.Style
	Section        lipgloss.Style
	Header         lipgloss.Style
	Content        lipgloss.Style
	StatusBar      lipgloss.Style
	StatusBarLeft  lipgloss.Style
	StatusBarRight lipgloss.Style
	Help           lipgloss.Style
	Key            lipgloss.Style
	Subtle         lipgloss.Style
	Error          lipgloss.Style
	Success        lipgloss.Style
	Warning        lipgloss.Style
	Info           lipgloss.Style
	Box            lipgloss.Style
	Prompt         lipgloss.Style
	Cursor         lipgloss.Style
	Selected       lipgloss.Style
	Unselected     lipgloss.Style
}

func New() *Styles {
	s := &Styles{}
	
	primary := lipgloss.Color("#7D56F4")
	secondary := lipgloss.Color("#F25D94")
	accent := lipgloss.Color("#FAFAFA")
	subtle := lipgloss.Color("#626262")
	success := lipgloss.Color("#04B575")
	warning := lipgloss.Color("#FFAA00")
	error := lipgloss.Color("#FF5555")
	info := lipgloss.Color("#00B4D8")
	
	s.Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(primary).
		MarginBottom(1)
	
	s.Subtitle = lipgloss.NewStyle().
		Foreground(secondary).
		Italic(true)
	
	s.Section = lipgloss.NewStyle().
		Bold(true).
		Foreground(primary).
		Underline(true).
		MarginBottom(1)
	
	s.Header = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(subtle).
		MarginBottom(1).
		Padding(1, 0).
		Align(lipgloss.Center)
	
	s.Content = lipgloss.NewStyle().
		Padding(1, 2).
		Align(lipgloss.Left)
	
	s.StatusBar = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#3C3C3C"))
	
	s.StatusBarLeft = lipgloss.NewStyle().
		Foreground(accent).
		Background(primary).
		Padding(0, 1)
	
	s.StatusBarRight = lipgloss.NewStyle().
		Foreground(accent).
		Background(secondary).
		Padding(0, 1)
	
	s.Help = lipgloss.NewStyle().
		Foreground(subtle)
	
	s.Key = lipgloss.NewStyle().
		Foreground(primary).
		Bold(true)
	
	s.Subtle = lipgloss.NewStyle().
		Foreground(subtle)
	
	s.Error = lipgloss.NewStyle().
		Foreground(error).
		Bold(true)
	
	s.Success = lipgloss.NewStyle().
		Foreground(success).
		Bold(true)
	
	s.Warning = lipgloss.NewStyle().
		Foreground(warning).
		Bold(true)
	
	s.Info = lipgloss.NewStyle().
		Foreground(info)
	
	s.Box = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(primary).
		Padding(1, 2)
	
	s.Prompt = lipgloss.NewStyle().
		Foreground(primary).
		Bold(true)
	
	s.Cursor = lipgloss.NewStyle().
		Background(primary).
		Foreground(accent)
	
	s.Selected = lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		PaddingLeft(2)
	
	s.Unselected = lipgloss.NewStyle().
		Foreground(subtle).
		PaddingLeft(4)
	
	return s
}

func (s *Styles) List(items []string, selected int) string {
	var result []string
	for i, item := range items {
		if i == selected {
			result = append(result, s.Selected.Render("▸ "+item))
		} else {
			result = append(result, s.Unselected.Render(item))
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, result...)
}

func (s *Styles) Table(headers []string, rows [][]string) string {
	table := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(s.Subtle.GetForeground())
	
	var tableRows []string
	
	headerRow := ""
	for _, h := range headers {
		headerRow += s.Title.Render(h) + "\t"
	}
	tableRows = append(tableRows, headerRow)
	
	for _, row := range rows {
		rowStr := ""
		for _, cell := range row {
			rowStr += cell + "\t"
		}
		tableRows = append(tableRows, rowStr)
	}
	
	return table.Render(lipgloss.JoinVertical(lipgloss.Left, tableRows...))
}