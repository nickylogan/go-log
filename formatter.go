package log

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/logrusorgru/aurora/v3"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	defaultSourceLine      = "<???>:<???>"
	defaultFunc            = "<???>"
	defaultTimestampFormat = time.RFC3339
)

type textFormatter struct {
	tty bool
	sync.Once
}

func (f *textFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = new(bytes.Buffer)
	}

	f.Do(func() { f.init(entry) })

	sourceLine, funcName := getCaller()
	shouldFormat := f.tty
	if shouldFormat {
		f.writeFormatted(b, entry, sourceLine, funcName, defaultScheme)
	} else {
		f.writePlain(b, entry, sourceLine, funcName)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *textFormatter) init(entry *logrus.Entry) {
	if entry.Logger == nil {
		return
	}

	file, ok := entry.Logger.Out.(*os.File)
	if !ok {
		return
	}
	f.tty = terminal.IsTerminal(int(file.Fd()))
}

func (f *textFormatter) writeFormatted(b *bytes.Buffer, entry *logrus.Entry, sourceLine, funcName string, scheme colorScheme) {
	levelColor := deriveLevelColor(entry.Level, scheme)
	levelText := fmt.Sprintf("%5s", getLevelText(entry.Level))

	timestamp := fmt.Sprintf("[%s]", entry.Time.Format(defaultTimestampFormat))

	timestamp = aurora.Colorize(timestamp, scheme.timestamp).String()
	levelText = aurora.Colorize(levelText, levelColor).String()
	sourceLine = aurora.Colorize(sourceLine, scheme.sourceLine).String()
	funcName = aurora.Colorize(funcName, scheme.funcName).String()

	// print prefix
	fmt.Fprintf(b, "%s %s %s %s: %s", timestamp, sourceLine, levelText, funcName, entry.Message)

	for k, v := range entry.Data {
		k := aurora.Colorize(k, levelColor)
		fmt.Fprintf(b, " %s=%+v", k, v)
	}
}

func (f *textFormatter) writePlain(b *bytes.Buffer, entry *logrus.Entry, sourceLine, funcName string) {
	f.writeKeyValue(b, "time", entry.Time.Format(defaultTimestampFormat))
	b.WriteByte(' ')
	f.writeKeyValue(b, "level", entry.Level.String())
	if entry.Message != "" {
		b.WriteByte(' ')
		f.writeKeyValue(b, "msg", entry.Message)
	}
	if sourceLine != defaultSourceLine {
		b.WriteByte(' ')
		f.writeKeyValue(b, "sourceLine", sourceLine)
	}

	for k, v := range entry.Data {
		b.WriteByte(' ')
		f.writeKeyValue(b, k, v)
	}
}

func (f *textFormatter) writeKeyValue(b *bytes.Buffer, key string, value interface{}) {
	b.WriteString(key)
	b.WriteByte('=')
	f.writeValue(b, value)
}

func (f *textFormatter) writeValue(b *bytes.Buffer, value interface{}) {
	switch v := value.(type) {
	case string:
		if f.shouldQuote(v) {
			fmt.Fprintf(b, "'%s'", v)
		} else {
			fmt.Fprintf(b, "%s", v)
		}
	case error:
		errText := v.Error()
		if f.shouldQuote(errText) {
			fmt.Fprintf(b, "'%s'", errText)
		} else {
			fmt.Fprintf(b, "%s", errText)
		}
	default:
		fmt.Fprint(b, value)
	}
}

func (f *textFormatter) shouldQuote(text string) bool {
	if text == "" {
		return true
	}

	for _, ch := range text {
		if !(('a' <= ch && ch <= 'z') ||
			('A' <= ch && ch <= 'Z') ||
			('0' <= ch && ch <= '9') ||
			ch == '.' ||
			ch == '-') {
			return true
		}
	}

	return false
}

func deriveLevelColor(level logrus.Level, scheme colorScheme) (color aurora.Color) {
	switch level {
	case logrus.PanicLevel:
		color = scheme.panicLevel
	case logrus.FatalLevel:
		color = scheme.fatalLevel
	case logrus.ErrorLevel:
		color = scheme.errorLevel
	case logrus.WarnLevel:
		color = scheme.warnLevel
	case logrus.InfoLevel:
		color = scheme.infoLevel
	default:
		color = scheme.debugLevel
	}

	return color
}

func getLevelText(level logrus.Level) string {
	var levelText string
	if level == logrus.WarnLevel {
		levelText = "warn"
	} else {
		levelText = level.String()
	}

	return strings.ToUpper(levelText)
}

func getCaller() (sourceLine, funcName string) {
	pc, file, line, ok := runtime.Caller(9)
	if !ok {
		return defaultSourceLine, defaultFunc
	}
	sourceLine = file + ":" + strconv.Itoa(line)

	funcInfo := runtime.FuncForPC(pc)
	if funcInfo == nil {
		funcName = defaultFunc
	} else {
		funcName = funcInfo.Name()
		idx := strings.LastIndex(funcName, ".")
		funcName = funcName[idx+1:] + "()"
	}

	return sourceLine, funcName
}
