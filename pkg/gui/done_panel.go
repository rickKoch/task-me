package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)


func (gui *Gui) refreshDone() error {
	gui.Log.Info("refreshDone")
	doneView := gui.getDoneView()
	if doneView == nil {
		return nil
	}


	if gui.State.Panels.Done.SelectedLine == -1 {
		gui.State.Panels.Done.SelectedLine = 0
	}

	if 1 < gui.State.Panels.Done.SelectedLine {
		gui.State.Panels.Done.SelectedLine = 1
	}

	gui.g.Update(func(g *gocui.Gui) error {
		doneView := gui.getDoneView()
		doneView.Clear()

		list := "done1\ndone2\n"
		fmt.Fprint(doneView, list)

		return nil
	})
	return nil
}
