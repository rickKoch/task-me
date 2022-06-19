package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) refreshTodo() error {
	gui.Log.Info("refreshTodo")
	todoView := gui.getTodoView()
	if todoView == nil {
		return nil
	}

	if gui.State.Panels.Todo.SelectedLine == -1 {
		gui.State.Panels.Todo.SelectedLine = 0
	}

	if 1 < gui.State.Panels.Todo.SelectedLine {
		gui.State.Panels.Todo.SelectedLine = 1
	}

	gui.g.Update(func(g *gocui.Gui) error {
		todoView := gui.getTodoView()
		todoView.Clear()

		//isFocused := gui.g.CurrentView().Name() == "todo"

		list := "todo1\ntodo2\n"
		fmt.Fprint(todoView, list)

		return nil
	})
	return nil
}

func (gui *Gui) handleTodoPrevLine(g *gocui.Gui, v *gocui.View) error {
	if gui.g.CurrentView() != v {
		return nil
	}

	panelState := gui.State.Panels.Todo
	gui.changeSelectedLine(&panelState.SelectedLine, 2, true)

	return gui.handleTodoSelect(gui.g, v)
}

func (gui *Gui) handleTodoNextLine(g *gocui.Gui, v *gocui.View) error {
	if gui.g.CurrentView() != v {
		return nil
	}

	panelState := gui.State.Panels.Todo
	gui.changeSelectedLine(&panelState.SelectedLine, 2, false)

	return gui.handleTodoSelect(gui.g, v)
}

func (gui *Gui) handleTodoSelect(g *gocui.Gui, v *gocui.View) error {
	gui.focusPoint(0, gui.State.Panels.Todo.SelectedLine, 2, v)
	return nil
}
