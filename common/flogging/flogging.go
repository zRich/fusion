package flogging

import (
	"fmt"
	"io"
)

// FormatLogger is a logger interface that output logs with a format
type FormatLogger interface {
	Tracef(format string, v ...any)
	Debugf(format string, v ...any)
	Infof(format string, v ...any)
	Noticef(format string, v ...any)
	Warnf(format string, v ...any)
	Errorf(format string, v ...any)
	Fatalf(format string, v ...any)
}

// Logger is a logger interface that output logs with levels
type Logger interface {
	Trace(v ...any)
	Debug(v ...any)
	Info(v ...any)
	Notice(v ...any)
	Warn(v ...any)
	Error(v ...any)
	Fatal(v ...any)
}

//config a logger
type Control interface {
	SetLevel(Level)
	SetOutput(io.Writer)
}

//combination of Logger, FormatLogger and Control
type FullLogger interface {
	Logger
	FormatLogger
	Control
}

//level defines log message
//when level is set,any log message with
//a lower log level will not be output
type Level int

//the levels of logs
const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelNotice
	LevelWarn
	LevelError
	LevelFatal
)

var strs = []string{
	"[Trace] ",
	"[Debug] ",
	"[Info] ",
	"[Notice] ",
	"[Warn] ",
	"[Error] ",
	"[Fatal] ",
}

func (lv Level) toString() string {
	if lv >= LevelTrace && lv <= LevelFatal {
		return strs[lv]
	}
	return fmt.Sprintf("[?%d] ", lv)
}
