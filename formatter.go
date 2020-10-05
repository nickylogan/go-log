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

var (
	// qualified package name, cached at first use
	gologPackage string

	// Positions in the call stack when tracing to report the calling method
	minCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

const (
	maxCallerDepth   int    = 25
	knownGologFrames int    = 8
	logrusPackage    string = "github.com/sirupsen/logrus"
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
	f.writeKeyValue(b, "level", entry.Level.String())
	if entry.Message != "" {
		f.writeKeyValue(b, "msg", entry.Message)
	}
	if sourceLine != defaultSourceLine {
		f.writeKeyValue(b, "sourceLine", sourceLine)
	}

	for k, v := range entry.Data {
		f.writeKeyValue(b, k, v)
	}
}

func (f *textFormatter) writeKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
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
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
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
	frame := getCallerFrame()
	if frame == nil {
		return defaultSourceLine, defaultFunc
	}
	sourceLine = frame.File + ":" + strconv.Itoa(frame.Line)

	if frame.Func == nil {
		funcName = defaultFunc
	} else {
		funcName = frame.Func.Name()
		idx := strings.LastIndex(funcName, ".")
		funcName = funcName[idx+1:] + "()"
	}

	return sourceLine, funcName
}

// This implementation is blatantly ripped-off from logrus. We can't use their
// ReportCaller flag because of us using different stack frames.
//
// Refer to https://github.com/sirupsen/logrus/blob/master/entry.go#L173.
func getCallerFrame() *runtime.Frame {
	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maxCallerDepth)
		_ = runtime.Callers(0, pcs)

		// dynamic get the package name and the minimum caller depth
		for i := 0; i < maxCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "getCallerFrame") {
				gologPackage = getPackageName(funcName)
				break
			}
		}

		minCallerDepth = knownGologFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != gologPackage && pkg != logrusPackage {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
