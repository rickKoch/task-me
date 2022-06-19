package main

import (
	"log"
	"os"

	"github.com/rickKoch/task-me/pkg/app"
	"github.com/rickKoch/task-me/pkg/config"
)

func main() {
  appConfig, err := config.NewAppConfig(true)
  if err != nil {
    log.Fatal(err.Error())
  }

	app, err := app.NewApp(appConfig)
	if err == nil {
		err = app.Run()
	}

	log.Println(err)
	os.Exit(0)
}
