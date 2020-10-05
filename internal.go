package log

import "github.com/sirupsen/logrus"

type extendedLogger interface {
	logrus.FieldLogger

	// Trace logs a message at level Trace.
	Trace(args ...interface{})
	// Tracef logs a message at level Trace.
	Tracef(format string, args ...interface{})
	// Traceln logs a message at level Trace.
	Traceln(args ...interface{})
}

type logger struct {
	log extendedLogger
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

func (l *logger) WithError(err error) Logger {
	return &logger{log: l.log.WithError(err)}
}

func (l *logger) WithField(key string, value interface{}) Logger {
	return &logger{log: l.log.WithField(key, value)}
}

func (l *logger) WithFields(fields Fields) Logger {
	return &logger{log: l.log.WithFields(logrus.Fields(fields))}
}
