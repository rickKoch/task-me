package gui

import (
	"sync"
	"time"

	throttle "github.com/boz/go-throttle"
	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"
)

type Gui struct {
	g             *gocui.Gui
	Log           *logrus.Entry
	waitForIntro  sync.WaitGroup
	CyclableViews []string
	State         guiState
}

type todoPanelState struct {
	SelectedLine int
}

type donePanelState struct {
	SelectedLine int
}

type cancelledPanelState struct {
	SelectedLine int
}

type panelStates struct {
	Todo      *todoPanelState
	Done      *donePanelState
	Cancelled *cancelledPanelState
}

type guiState struct {
	Panels *panelStates
}

func NewGui(log *logrus.Entry) (*Gui, error) {
	initialState := guiState{
		Panels: &panelStates{
			Todo:      &todoPanelState{SelectedLine: -1},
			Done:      &donePanelState{SelectedLine: -1},
			Cancelled: &cancelledPanelState{SelectedLine: -1},
		},
	}

	gui := &Gui{
		Log:           log,
		CyclableViews: []string{"project", "todo", "done", "cancelled"},
		State:         initialState,
	}

	return gui, nil
}

func (gui *Gui) Run() error {
	g, err := gocui.NewGui(gocui.Output256)

	if err != nil {
		return err
	}
	defer g.Close()

	gui.g = g

	if err := gui.SetColorScheme(); err != nil {
		return nil
	}

	gui.waitForIntro.Add(1)

	throttledRefresh := throttle.ThrottleFunc(time.Microsecond*50, true, gui.refresh)
	defer throttledRefresh.Stop()

	go func() {
		gui.Log.Info("before->waitForIntro")
		gui.waitForIntro.Wait()
		gui.Log.Info("after->waitForIntro")
		throttledRefresh.Trigger()
	}()

	g.SetManager(gocui.ManagerFunc(gui.layout), gui.getFocusLayout())

	if err := gui.keybindings(g); err != nil {
		return err
	}

	err = g.MainLoop()

	return err
}

func (gui *Gui) refresh() {
	go gui.refreshProject()
	go func() {
		if err := gui.refreshTodo(); err != nil {
			gui.Log.Error(err)
		}
	}()

	go func() {
		if err := gui.refreshDone(); err != nil {
			gui.Log.Error(err)
		}
	}()

	go func() {
		if err := gui.refreshCancelled(); err != nil {
			gui.Log.Error(err)
		}
	}()
}

func (gui *Gui) quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}
