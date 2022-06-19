package gui

import (
	"github.com/jroimartin/gocui"
)

func (gui *Gui) layout(g *gocui.Gui) error {
	g.Highlight = true
	width, height := g.Size()

	minimumHeight := 9
	minimumWidth := 10

	if height < minimumHeight || width < minimumWidth {
		v, err := g.SetView("limit", 0, 0, width-1, height-1)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}

			v.Title = "not enough spce"
			v.Wrap = true
			_, _ = g.SetCurrentView("limit")
		}

		return nil
	}

	usableSpace := height - 1
	tallPanels := 3

	vHeights := map[string]int{
		"project":   3,
		"todo":      usableSpace/tallPanels + usableSpace%tallPanels,
		"done":      usableSpace/tallPanels + usableSpace%tallPanels + usableSpace/tallPanels,
		"cancelled": usableSpace/tallPanels + usableSpace%tallPanels + usableSpace/tallPanels + usableSpace/tallPanels,
		"options":   1,
	}

	leftSideWidth := width / 3

	v, err := g.SetView("main", leftSideWidth+1, 0, width-1, height-2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.FgColor = gocui.ColorDefault
	}

	if v, err := g.SetView("project", 0, 0, leftSideWidth, vHeights["project"]-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Project"
		v.FgColor = gocui.ColorDefault
	}

	todoView, err := g.SetView("todo", 0, vHeights["project"], leftSideWidth, vHeights["todo"]-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		todoView.Highlight = true
		todoView.Title = "Todos"
		todoView.FgColor = gocui.ColorDefault
	}

	doneView, err := g.SetView("done", 0, vHeights["todo"], leftSideWidth, vHeights["done"]-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		doneView.Highlight = true
		doneView.Title = "Done"
		doneView.FgColor = gocui.ColorDefault
	}

	cancelledView, err := g.SetView("cancelled", 0, vHeights["done"], leftSideWidth, vHeights["cancelled"]-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		cancelledView.Highlight = true
		cancelledView.Title = "Cancelled"
		cancelledView.FgColor = gocui.ColorDefault

		// move this from here
		gui.waitForIntro.Done()
	}

	if gui.g.CurrentView() == nil {
		v, err := gui.g.View("project")
		if err != nil {
			return err
		}

		if err := gui.switchFocus(gui.g, nil, v, false); err != nil {
			return err
		}
	}

	return nil
}

func (gui *Gui) getFocusLayout() gocui.ManagerFunc {
	var previousView *gocui.View
	return func(g *gocui.Gui) error {
		if err := gui.onFocusChange(); err != nil {
			return err
		}

		newView := gui.g.CurrentView()
		if newView != previousView && !gui.isPopupPanel(newView.Name()) {
			gui.onFocusLost(previousView, newView)
			gui.onFocus(newView)
			previousView = newView
		}

		return nil
	}
}

func (gui *Gui) onFocusChange() error {
	currentView := gui.g.CurrentView()
	for _, view := range gui.g.Views() {
		view.Highlight = view == currentView && view.Name() != "main"
	}
	return nil
}

func (gui *Gui) onFocusLost(v *gocui.View, newView *gocui.View) {
	if v == nil {
		return
	}

	// refocusing because in responsive mode (when the window is very short) we
	// want to ensure that after the view size changes we can still see the last
	// selected item
	gui.focusPointInView(v)

	gui.Log.Info(v.Name() + " focus lost")
}

func (gui *Gui) onFocus(v *gocui.View) {
	if v == nil {
		return
	}

	gui.focusPointInView(v)

	gui.Log.Info(v.Name() + " focus gained")
}

type listViewState struct {
	selectedLine int
	lineCount    int
}

func (gui *Gui) focusPointInView(view *gocui.View) {
	if view == nil {
		return
	}

	listViews := map[string]listViewState{
		"todo": {selectedLine: gui.State.Panels.Todo.SelectedLine, lineCount: 2},
	}

	if state, ok := listViews[view.Name()]; ok {
		gui.focusPoint(0, state.selectedLine, state.lineCount, view)
	}
}
