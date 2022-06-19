package gui

import "github.com/jroimartin/gocui"

type Binding struct {
	ViewName    string
	Handler     func(*gocui.Gui, *gocui.View) error
	Key         interface{}
	Modifier    gocui.Modifier
	Description string
}

func (gui *Gui) GetInitialKeybindings() []*Binding {
	bindings := []*Binding{
		{
			ViewName: "",
			Key:      'q',
			Modifier: gocui.ModNone,
			Handler:  gui.quit,
		},
		{
			ViewName: "",
			Key:      gocui.KeyCtrlC,
			Modifier: gocui.ModNone,
			Handler:  gui.quit,
		},
		{
			ViewName: "",
			Key:      gocui.KeyEsc,
			Modifier: gocui.ModNone,
			Handler:  gui.quit,
		},
	}

	for _, viewName := range []string{"project", "todo", "done", "cancelled"} {
		bindings = append(bindings, []*Binding{
			{ViewName: viewName, Key: gocui.KeyArrowLeft, Modifier: gocui.ModNone, Handler: gui.previousView},
			{ViewName: viewName, Key: gocui.KeyArrowRight, Modifier: gocui.ModNone, Handler: gui.nextView},
			{ViewName: viewName, Key: 'h', Modifier: gocui.ModNone, Handler: gui.previousView},
			{ViewName: viewName, Key: 'l', Modifier: gocui.ModNone, Handler: gui.nextView},
		}...)
	}

	panelMap := map[string]struct {
		onKeyUpPress   func(*gocui.Gui, *gocui.View) error
		onKeyDownPress func(*gocui.Gui, *gocui.View) error
	}{
		"todo": {onKeyUpPress: gui.handleTodoPrevLine, onKeyDownPress: gui.handleTodoNextLine},
		"done": {onKeyUpPress: gui.handleTodoPrevLine, onKeyDownPress: gui.handleTodoNextLine},
		"cancelled": {onKeyUpPress: gui.handleTodoPrevLine, onKeyDownPress: gui.handleTodoNextLine},
	}

	for viewName, functions := range panelMap {
		bindings = append(bindings, []*Binding{
			{ViewName: viewName, Key: 'k', Modifier: gocui.ModNone, Handler: functions.onKeyUpPress},
			{ViewName: viewName, Key: gocui.KeyArrowUp, Modifier: gocui.ModNone, Handler: functions.onKeyUpPress},
			{ViewName: viewName, Key: gocui.MouseWheelUp, Modifier: gocui.ModNone, Handler: functions.onKeyUpPress},
			{ViewName: viewName, Key: 'j', Modifier: gocui.ModNone, Handler: functions.onKeyDownPress},
			{ViewName: viewName, Key: gocui.KeyArrowDown, Modifier: gocui.ModNone, Handler: functions.onKeyDownPress},
			{ViewName: viewName, Key: gocui.MouseWheelDown, Modifier: gocui.ModNone, Handler: functions.onKeyDownPress},
		}...)
	}

	return bindings
}

func (gui *Gui) keybindings(g *gocui.Gui) error {
	bindings := gui.GetInitialKeybindings()

	for _, binding := range bindings {
		if err := g.SetKeybinding(binding.ViewName, binding.Key, binding.Modifier, binding.Handler); err != nil {
			return err
		}
	}

	return nil
}
