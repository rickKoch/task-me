package app

import (
	"github.com/rickKoch/task-me/pkg/config"
	"github.com/rickKoch/task-me/pkg/gui"
	"github.com/rickKoch/task-me/pkg/log"
	"github.com/sirupsen/logrus"
)

type App struct {
	Config *config.AppConfig
	Gui    *gui.Gui
	Log    *logrus.Entry
}

func NewApp(config *config.AppConfig) (*App, error) {
	app := &App{
		Config:  config,
	}
	var err error

	app.Log = log.NewLogger(app.Config)
	app.Gui, err = gui.NewGui(app.Log)
	if err != nil {
		return app, err
	}

	return app, nil
}

func (app *App) Run() error {
	return app.Gui.RunWithSubprocesses()
}
