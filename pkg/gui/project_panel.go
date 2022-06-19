package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) refreshProject() {
  gui.Log.Info("refreshProject")
  v := gui.getProjectView()

  projectName := gui.getProjectName()

  gui.g.Update(func(*gocui.Gui) error {
    v.Clear()
    fmt.Fprint(v, projectName)
    return nil
  })
}

func (gui *Gui) getProjectName() string {
  return "test project"
}
