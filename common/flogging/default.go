package flogging

import (
	"fmt"
	"io"
	"log"
	"os"
)

var logger FullLogger = &defaultLogger{
	stdlog: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
}

//SetOutput sets the output of default logger
func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

//SetLevel sets the level
//default log level is LevelTrace
func SetLevel(lv Level) {
	logger.SetLevel(lv)
}

//DefaultLogger return the default
func DefaultLogger() FullLogger {
	return logger
}

//SetLogger sets the default loggerr
func SetLogger(v FullLogger) {
	logger = v
}

//Fatal calls the default logger's Fatal method and then os.Exit(1)
func Fatal(v ...any) {
	logger.Fatal(v...)
}

//Error calls the default logger's Error method
func Error(v ...any) {
	logger.Error(v...)
}

//Warn calls the default logger's Warn method
func Warn(v ...any) {
	logger.Warn(v...)
}

//Notice calls the default logger's Notice method
func Notice(v ...any) {
	logger.Notice(v...)
}

//Info calls the default logger's Info method
func Info(v ...any) {
	logger.Info(v...)
}

//Debug calls the default logger's Debug method
func Debug(v ...any) {
	logger.Debug(v...)
}

//Trace calls the default logger's Trace method
func Trace(v ...any) {
	logger.Trace(v...)
}

//Fatalf calls the default logger's Fatalf method and then os.Exit(1)
func Fatalf(format string, v ...any) {
	logger.Fatalf(format, v...)
}

//Errorf calls the default logger's Errorf method
func Errorf(format string, v ...any) {
	logger.Errorf(format, v...)
}

//Warnf calls the default logger's Warnf method
func Warnf(format string, v ...any) {
	logger.Warnf(format, v...)
}

//Noticef calls the default logger's Noticef method
func Noticef(format string, v ...any) {
	logger.Noticef(format, v...)
}

//Infof calls the default logger's Infof method
func Infof(format string, v ...any) {
	logger.Infof(format, v...)
}

//Debugf calls the default logger's Debugf method
func Debugf(format string, v ...any) {
	logger.Debugf(format, v...)
}

//Tracef calls the default logger's Tracef method
func Tracef(format string, v ...any) {
	logger.Tracef(format, v...)
}

type defaultLogger struct {
	stdlog *log.Logger
	level  Level
}

func (ll *defaultLogger) SetOutput(w io.Writer) {
	ll.stdlog.SetOutput(w)
}

func (ll *defaultLogger) SetLevel(lv Level) {
	ll.level = lv
}

func (ll *defaultLogger) logf(lv Level, format *string, v ...any) {
	if ll.level > lv {
		return
	}
	msg := lv.toString()
	if format != nil {
		msg += fmt.Sprintf(*format, v...)
	} else {
		msg += fmt.Sprint(v...)
	}
	ll.stdlog.Output(4, msg)
	if lv == LevelFatal {
		os.Exit(1)
	}
}

func (ll *defaultLogger) Fatal(v ...any) {
	ll.logf(LevelFatal, nil, v...)
}

func (ll *defaultLogger) Error(v ...any) {
	ll.logf(LevelError, nil, v...)
}

func (ll *defaultLogger) Warn(v ...any) {
	ll.logf(LevelWarn, nil, v...)
}
func (ll *defaultLogger) Notice(v ...any) {
	ll.logf(LevelNotice, nil, v...)
}
func (ll *defaultLogger) Info(v ...any) {
	ll.logf(LevelInfo, nil, v...)
}
func (ll *defaultLogger) Debug(v ...any) {
	ll.logf(LevelDebug, nil, v...)
}
func (ll *defaultLogger) Trace(v ...any) {
	ll.logf(LevelTrace, nil, v...)
}

func (ll *defaultLogger) Fatalf(format string, v ...any) {
	ll.logf(LevelFatal, &format, v...)
}

func (ll *defaultLogger) Errorf(format string, v ...any) {
	ll.logf(LevelError, &format, v...)
}

func (ll *defaultLogger) Warnf(format string, v ...any) {
	ll.logf(LevelWarn, &format, v...)
}
func (ll *defaultLogger) Noticef(format string, v ...any) {
	ll.logf(LevelNotice, &format, v...)
}
func (ll *defaultLogger) Infof(format string, v ...any) {
	ll.logf(LevelInfo, &format, v...)
}
func (ll *defaultLogger) Debugf(format string, v ...any) {
	ll.logf(LevelDebug, &format, v...)
}
func (ll *defaultLogger) Tracef(format string, v ...any) {
	ll.logf(LevelTrace, &format, v...)
}
