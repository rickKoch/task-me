package gui

import "github.com/jroimartin/gocui"

func (gui *Gui) nextView(g *gocui.Gui, v *gocui.View) error {
	var focusedViewName string

	gui.Log.Info("VIEW_NAME::", v.Name())
	gui.Log.Info(v.Name() == gui.CyclableViews[len(gui.CyclableViews)-1])

	if v == nil || v.Name() == gui.CyclableViews[len(gui.CyclableViews)-1] {
		focusedViewName = gui.CyclableViews[0]
	} else {
		viewName := v.Name()
		for i := range gui.CyclableViews {
			if viewName == gui.CyclableViews[i] {
				focusedViewName = gui.CyclableViews[i+1]
				break
			}

			if i == len(gui.CyclableViews)-1 {
				gui.Log.Info("not in list of views")
				return nil
			}
		}
	}

	focusedView, err := g.View(focusedViewName)
	if err != nil {
		panic(err)
	}

	return gui.switchFocus(g, v, focusedView, false)
}

func (gui *Gui) previousView(g *gocui.Gui, v *gocui.View) error {
	var focusedViewName string
	if v == nil || v.Name() == gui.CyclableViews[0] {
		focusedViewName = gui.CyclableViews[len(gui.CyclableViews)-1]
	} else {
		viewName := v.Name()
		for i := range gui.CyclableViews {
			if viewName == gui.CyclableViews[i] {
				focusedViewName = gui.CyclableViews[i-1]
				break
			}
			if i == len(gui.CyclableViews)-1 {
				gui.Log.Info("not in list of views")
				return nil
			}
		}
	}
	focusedView, err := g.View(focusedViewName)
	if err != nil {
		panic(err)
	}
	return gui.switchFocus(g, v, focusedView, false)
}

func (gui *Gui) changeSelectedLine(line *int, total int, up bool) {
	if up {
		if *line == -1 || *line == 0 {
			return
		}

		*line -= 1
	} else {
		if *line == -1 || *line == total-1 {
			return
		}

		*line += 1
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (gui *Gui) isPopupPanel(viewName string) bool {
	return false
}

func (gui *Gui) focusPoint(selectedX int, selectedY int, lineCount int, v *gocui.View) {
	if selectedY < 0 || selectedY > lineCount {
		return
	}
	ox, oy := v.Origin()
	originalOy := oy
	cx, cy := v.Cursor()
	originalCy := cy
	_, height := v.Size()

	ly := Max(height-1, 0)

	windowStart := oy
	windowEnd := oy + ly

	if selectedY < windowStart {
		oy = Max(oy-(windowStart-selectedY), 0)
	} else if selectedY > windowEnd {
		oy += (selectedY - windowEnd)
	}

	if windowEnd > lineCount-1 {
		shiftAmount := (windowEnd - (lineCount - 1))
		oy = Max(oy-shiftAmount, 0)
	}

	if originalOy != oy {
		_ = v.SetOrigin(ox, oy)
	}

	cy = selectedY - oy
	if originalCy != cy {
		_ = v.SetCursor(cx, selectedY-oy)
	}
}

func (gui *Gui) switchFocus(g *gocui.Gui, oldView, newView *gocui.View, returning bool) error {
	//if oldView != nil && !gui.isPopupPanel(oldView.Name()) && !returning {

	//}

	gui.Log.Info("setting highlight to true for view " + newView.Name())
	gui.Log.Info("new focused view is " + newView.Name())

	if _, err := g.SetCurrentView(newView.Name()); err != nil {
		return err
	}

	if _, err := g.SetViewOnTop(newView.Name()); err != nil {
		return err
	}

	return nil
}

func (gui *Gui) getProjectView() *gocui.View {
	v, _ := gui.g.View("project")
	return v
}

func (gui *Gui) getTodoView() *gocui.View {
	v, _ := gui.g.View("todo")
	return v
}

func (gui *Gui) getDoneView() *gocui.View {
	v, _ := gui.g.View("done")
	return v
}

func (gui *Gui) getCancelledView() *gocui.View {
	v, _ := gui.g.View("cancelled")
	return v
}

