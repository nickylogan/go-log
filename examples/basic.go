package main

import (
	"errors"

	log "github.com/nickylogan/go-log"
)

func main() {
	defer func() {
		_ = recover()
	}()

	log.Init(log.WithLevel(log.TraceLevel))
	log.WithError(errors.New("trace error")).
		WithFields(log.Fields{"field": "trace"}).
		Trace("trace message")
	log.WithError(errors.New("debug error")).
		WithFields(log.Fields{"field": "debug"}).
		Debug("debug message")
	log.WithError(errors.New("info error")).
		WithFields(log.Fields{"field": "info"}).
		Info("info message")
	log.WithError(errors.New("warn error")).
		WithFields(log.Fields{"field": "warn"}).
		Warn("warn message")
	log.WithError(errors.New("error error")).
		WithFields(log.Fields{"field": "error"}).
		Error("error message")

	log.Trace("trace message")
	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")
}
