package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Option configures the logger.
type Option interface {
	apply(*logrus.Logger)
}

type optionFunc func(*logrus.Logger)

func (f optionFunc) apply(logger *logrus.Logger) { f(logger) }

// WithLevel sets the logger level.
func WithLevel(level Level) Option {
	return optionFunc(func(logger *logrus.Logger) {
		logger.SetLevel(logrus.Level(level))
	})
}

// WithOutput sets the logger output
func WithOutput(writer io.Writer) Option {
	return optionFunc(func(logger *logrus.Logger) {
		logger.SetOutput(writer)
	})
}
