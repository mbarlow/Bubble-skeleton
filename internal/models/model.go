package models

import (
	"github.com/mbarlow/bubble-skeleton/internal/commands"
	"github.com/mbarlow/bubble-skeleton/internal/config"
	"github.com/mbarlow/bubble-skeleton/internal/views"
	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	StateMain State = iota
	StateHelp
	StateInput
)

type Model struct {
	config       *config.Config
	state        State
	width        int
	height       int
	ready        bool
	err          error
	inputValue   string
	cursor       int
	view         *views.View
}

func New(cfg *config.Config) Model {
	return Model{
		config: cfg,
		state:  StateMain,
		view:   views.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		commands.InitialLoad(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.view.SetSize(msg.Width, msg.Height)
		if !m.ready {
			m.ready = true
		}
		return m, nil

	case tea.KeyMsg:
		switch m.state {
		case StateMain:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "h", "?":
				m.state = StateHelp
			case "i":
				m.state = StateInput
				m.cursor = len(m.inputValue)
			case "r":
				cmds = append(cmds, commands.Refresh())
			}
		
		case StateHelp:
			switch msg.String() {
			case "ctrl+c", "q", "esc":
				m.state = StateMain
			}
		
		case StateInput:
			switch msg.String() {
			case "ctrl+c", "esc":
				m.state = StateMain
			case "enter":
				m.state = StateMain
				cmds = append(cmds, commands.ProcessInput(m.inputValue))
				m.inputValue = ""
				m.cursor = 0
			case "backspace":
				if m.cursor > 0 {
					m.inputValue = m.inputValue[:m.cursor-1] + m.inputValue[m.cursor:]
					m.cursor--
				}
			case "left":
				if m.cursor > 0 {
					m.cursor--
				}
			case "right":
				if m.cursor < len(m.inputValue) {
					m.cursor++
				}
			case "home":
				m.cursor = 0
			case "end":
				m.cursor = len(m.inputValue)
			default:
				if len(msg.String()) == 1 {
					m.inputValue = m.inputValue[:m.cursor] + msg.String() + m.inputValue[m.cursor:]
					m.cursor++
				}
			}
		}

	case commands.LoadCompleteMsg:
		m.ready = true
		
	case commands.RefreshCompleteMsg:
		
	case commands.ErrorMsg:
		m.err = msg.Err
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.ready {
		return m.view.Loading()
	}

	switch m.state {
	case StateHelp:
		return m.view.Help()
	case StateInput:
		return m.view.Input(m.inputValue, m.cursor)
	default:
		return m.view.Main()
	}
}