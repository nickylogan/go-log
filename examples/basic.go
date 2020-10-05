package main

import (
	log "github.com/nickylogan/go-log"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	log.Trace("trace message")
	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")
}
