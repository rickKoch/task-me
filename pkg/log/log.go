package log

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rickKoch/task-me/pkg/config"
	"github.com/sirupsen/logrus"
)

func NewLogger(config *config.AppConfig) *logrus.Entry {
	var log *logrus.Logger
	if config.Debug || os.Getenv("DEBUG") == "TRUE" {
		log = newDevelopomentLogger()
	} else {
		log = newProductionLogger()
	}

	log.Formatter = &logrus.JSONFormatter{}

	return log.WithFields(logrus.Fields{
		"debug": config.Debug,
	})
}

func newDevelopomentLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	file, err := os.OpenFile("./development.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Println("unable to log a file")
		os.Exit(1)
	}
	log.SetOutput(file)
	return log
}

func newProductionLogger() *logrus.Logger {
	log := logrus.New()
	log.Out = ioutil.Discard
	log.SetLevel(logrus.ErrorLevel)
	return log
}
