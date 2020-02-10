package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/tea"
	"github.com/charmbracelet/teaparty/spinner"
	"github.com/muesli/termenv"
)

var (
	color = termenv.ColorProfile()
)

type Model struct {
	spinner spinner.Model
}

func main() {
	p := tea.NewProgram(initialize, update, view, subscriptions)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func initialize() (tea.Model, tea.Cmd) {
	m := spinner.NewModel()
	m.Type = spinner.Dot

	return Model{
		spinner: m,
	}, nil
}

func update(msg tea.Msg, model tea.Model) (tea.Model, tea.Cmd) {
	m, ok := model.(Model)
	if !ok {
		return model, nil
	}

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			fallthrough
		case "esc":
			fallthrough
		case "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}

	default:
		m.spinner, _ = spinner.Update(msg, m.spinner)
		return m, nil
	}

}

func view(model tea.Model) string {
	m, ok := model.(Model)
	if !ok {
		return tea.ModelAssertionErr.String()
	}
	s := termenv.
		String(spinner.View(m.spinner)).
		Foreground(color.Color("205")).
		String()
	return fmt.Sprintf("\n\n   %s Loading forever...press q to quit\n\n", s)
}

func subscriptions(_ tea.Model) tea.Subs {
	return tea.Subs{
		"tick": func(model tea.Model) tea.Msg {
			m, ok := model.(Model)
			if !ok {
				return tea.ModelAssertionErr
			}
			return spinner.Sub(m.spinner)
		},
	}
}
