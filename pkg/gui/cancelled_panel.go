package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)


func (gui *Gui) refreshCancelled() error {
	gui.Log.Info("refreshCancelled")
	cancelledView := gui.getCancelledView()
	if cancelledView == nil {
		return nil
	}

	if gui.State.Panels.Cancelled.SelectedLine == -1 {
		gui.State.Panels.Cancelled.SelectedLine = 0
	}

	if 1 < gui.State.Panels.Cancelled.SelectedLine {
		gui.State.Panels.Cancelled.SelectedLine = 1
	}

	gui.g.Update(func(g *gocui.Gui) error {
		cancelledView := gui.getCancelledView()
		cancelledView.Clear()

		list := "cancelled1\ncancelled2\n"
		fmt.Fprint(cancelledView, list)

		return nil
	})
	return nil
}
