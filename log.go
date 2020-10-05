package log

// Logger defines a logger
type Logger interface {
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Traceln(args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})

	WithError(err error) Logger
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
}

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

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

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) Logger {
	return log().WithError(err)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) Logger {
	return log().WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) Logger {
	return log().WithFields(fields)
}
