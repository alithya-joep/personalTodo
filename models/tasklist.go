package models

import (
	"github.com/alithya-joep/personalTodo/dict"
	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	list list.Model
	err  error
}

func (m *Model) initList() {
	var todo dict.Status = dict.Todo

	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	m.list.Title = "Tasks"
	m.list.SetItems([]list.Item{
		dict.Task{Status: todo, Title: "devops", Description: "update due dates"},

		dict.Task{Status: todo, Title: "devops", Description: "update due dates"},
		dict.Task{Status: todo, Title: "timesheet", Description: "update oracle"},
	})
}
