package log

import (
	"github.com/sirupsen/logrus"
)

// Logger defines a logger
type Logger interface {
	// Trace logs a message at level Trace.
	Trace(args ...interface{})
	// Tracef logs a message at level Trace.
	Tracef(format string, args ...interface{})
	// Traceln logs a message at level Trace.
	Traceln(args ...interface{})

	// Debug logs a message at level Debug.
	Debug(args ...interface{})
	// Debugf logs a message at level Debug.
	Debugf(format string, args ...interface{})
	// Debugln logs a message at level Debug.
	Debugln(args ...interface{})

	// Info logs a message at level Info.
	Info(args ...interface{})
	// Infof logs a message at level Info.
	Infof(format string, args ...interface{})
	// Infoln logs a message at level Info.
	Infoln(args ...interface{})

	// Warn logs a message at level Warn.
	Warn(args ...interface{})
	// Warnf logs a message at level Warn.
	Warnf(format string, args ...interface{})
	// Warnln logs a message at level Warn.
	Warnln(args ...interface{})

	// Error logs a message at level Error.
	Error(args ...interface{})
	// Errorf logs a message at level Error.
	Errorf(format string, args ...interface{})
	// Errorln logs a message at level Error.
	Errorln(args ...interface{})

	// Fatal logs a message at level Fatal.log.
	Fatal(args ...interface{})
	// Fatalf logs a message at level Fatal.log.
	Fatalf(format string, args ...interface{})
	// Fatalln logs a message at level atal.log.
	Fatalln(args ...interface{})

	// Panic logs a message at level Panic.
	Panic(args ...interface{})
	// Panicf logs a message at level Panic.
	Panicf(format string, args ...interface{})
	// Panicln logs a message at level Panic.
	Panicln(args ...interface{})
}

// Trace logs a message at level Trace.
func Trace(args ...interface{}) {
	log().Trace(args...)
}

// Tracef logs a message at level Trace.
func Tracef(format string, args ...interface{}) {
	log().Tracef(format, args...)
}

// Traceln logs a message at level Trace.
func Traceln(args ...interface{}) {
	log().Traceln(args...)
}

// Debug logs a message at level Debug.
func Debug(args ...interface{}) {
	log().Debug(args...)
}

// Debugf logs a message at level Debug.
func Debugf(format string, args ...interface{}) {
	log().Debugf(format, args...)
}

// Debugln logs a message at level Debug.
func Debugln(args ...interface{}) {
	log().Debugln(args...)
}

// Info logs a message at level Info.
func Info(args ...interface{}) {
	log().Info(args...)
}

// Infof logs a message at level Info.
func Infof(format string, args ...interface{}) {
	log().Infof(format, args...)
}

// Infoln logs a message at level Info.
func Infoln(args ...interface{}) {
	log().Infoln(args...)
}

// Warn logs a message at level Warn.
func Warn(args ...interface{}) {
	log().Warn(args...)
}

// Warnf logs a message at level Warn.
func Warnf(format string, args ...interface{}) {
	log().Warnf(format, args...)
}

// Warnln logs a message at level Warn.
func Warnln(args ...interface{}) {
	log().Warnln(args...)
}

// Error logs a message at level Error.
func Error(args ...interface{}) {
	log().Error(args...)
}

// Errorf logs a message at level Error.
func Errorf(format string, args ...interface{}) {
	log().Errorf(format, args...)
}

// Errorln logs a message at level Error.
func Errorln(args ...interface{}) {
	log().Errorln(args...)
}

// Fatal logs a message at level Fatal.log.
func Fatal(args ...interface{}) {
	log().Fatal(args...)
}

// Fatalf logs a message at level Fatal.log.
func Fatalf(format string, args ...interface{}) {
	log().Fatalf(format, args...)
}

// Fatalln logs a message at level atal.log.
func Fatalln(args ...interface{}) {
	log().Fatalln(args...)
}

// Panic logs a message at level Panic.
func Panic(args ...interface{}) {
	log().Panic(args...)
}

// Panicf logs a message at level Panic.
func Panicf(format string, args ...interface{}) {
	log().Panicf(format, args...)
}

// Panicln logs a message at level Panic.
func Panicln(args ...interface{}) {
	log().Panicln(args...)
}

type logger struct {
	log Logger
}

var globalLogger *logger

func init() {
	iLog := logrus.StandardLogger()
	iLog.SetFormatter(&textFormatter{})
	globalLogger = &logger{log: iLog}
}

func log() Logger {
	return globalLogger
}

// Trace logs a message at level Trace.
func (l *logger) Trace(args ...interface{}) {
	l.log.Trace(args...)
}

// Tracef logs a message at level Trace.
func (l *logger) Tracef(format string, args ...interface{}) {
	l.log.Tracef(format, args...)
}

// Traceln logs a message at level Trace.
func (l *logger) Traceln(args ...interface{}) {
	l.log.Traceln(args...)
}

// Debug logs a message at level Debug.
func (l *logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

// Debugf logs a message at level Debug.
func (l *logger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

// Debugln logs a message at level Debug.
func (l *logger) Debugln(args ...interface{}) {
	l.log.Debugln(args...)
}

// Info logs a message at level Info.
func (l *logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

// Infof logs a message at level Info.
func (l *logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

// Infoln logs a message at level Info.
func (l *logger) Infoln(args ...interface{}) {
	l.log.Infoln(args...)
}

// Warn logs a message at level Warn.
func (l *logger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

// Warnf logs a message at level Warn.
func (l *logger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

// Warnln logs a message at level Warn.
func (l *logger) Warnln(args ...interface{}) {
	l.log.Warnln(args...)
}

// Error logs a message at level Error.
func (l *logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

// Errorf logs a message at level Error.
func (l *logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

// Errorln logs a message at level Error.
func (l *logger) Errorln(args ...interface{}) {
	l.log.Errorln(args...)
}

// Fatal logs a message at level Fatal.log.
func (l *logger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

// Fatalf logs a message at level Fatal.log.
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.log.Fatalf(format, args...)
}

// Fatalln logs a message at level atal.log.
func (l *logger) Fatalln(args ...interface{}) {
	l.log.Fatalln(args...)
}

// Panic logs a message at level Panic.
func (l *logger) Panic(args ...interface{}) {
	l.log.Panic(args...)
}

// Panicf logs a message at level Panic.
func (l *logger) Panicf(format string, args ...interface{}) {
	l.log.Panicf(format, args...)
}

// Panicln logs a message at level Panic.
func (l *logger) Panicln(args ...interface{}) {
	l.log.Panicln(args...)
}
