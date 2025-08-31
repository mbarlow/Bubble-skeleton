package commands

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type LoadCompleteMsg struct{}

type RefreshCompleteMsg struct {
	Data interface{}
}

type ErrorMsg struct {
	Err error
}

type ProcessInputMsg struct {
	Value string
}

func InitialLoad() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(500 * time.Millisecond)
		return LoadCompleteMsg{}
	}
}

func Refresh() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(200 * time.Millisecond)
		return RefreshCompleteMsg{
			Data: "Refreshed at " + time.Now().Format("15:04:05"),
		}
	}
}

func ProcessInput(value string) tea.Cmd {
	return func() tea.Msg {
		return ProcessInputMsg{
			Value: value,
		}
	}
}

func PerformAsyncOperation() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(1 * time.Second)
		return LoadCompleteMsg{}
	}
}

func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}