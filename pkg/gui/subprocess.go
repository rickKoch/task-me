package gui

import "github.com/jroimartin/gocui"


func (gui *Gui) RunWithSubprocesses() error {
  gui.Log.Info("RunWithSubprocesses")
  for {
    gui.Log.Info("RunWithSubprocesses->Loop")
    if err := gui.Run(); err != nil {
      if err == gocui.ErrQuit {
        break
      } else {
        return err
      }
    }
  }

  return nil
}
