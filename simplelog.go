package simplelog

/*
   potentially use: https://github.com/op/go-logging
*/

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   = log.New(os.Stdout, "TRC ", log.Ldate|log.Ltime)
	Debug   = log.New(os.Stdout, "DBG ", log.Ldate|log.Ltime)
	Info    = log.New(os.Stdout, "INF ", log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, "WRN ", log.Ldate|log.Ltime)
	Error   = log.New(os.Stderr, "ERR ", log.Ldate|log.Ltime)
	Fatal   = log.New(os.Stderr, "FTL ", log.Ldate|log.Ltime)
)

func resetLogLevel() {
	Trace = log.New(os.Stdout, "TRC ", log.Ldate|log.Ltime)
	Debug = log.New(os.Stdout, "DBG ", log.Ldate|log.Ltime)
	Info = log.New(os.Stdout, "INF ", log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, "WRN ", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "ERR ", log.Ldate|log.Ltime)
}

func SetLogLevel(level string) error {
	resetLogLevel()

	switch level {
	case "trace":
		break

	case "debug":
		Trace = log.New(ioutil.Discard, "TRC ", log.Ldate|log.Ltime)

	case "info":
		Trace = log.New(ioutil.Discard, "TRC ", log.Ldate|log.Ltime)
		Debug = log.New(ioutil.Discard, "DBG ", log.Ldate|log.Ltime)

	case "warning":
		Trace = log.New(ioutil.Discard, "TRC ", log.Ldate|log.Ltime)
		Debug = log.New(ioutil.Discard, "DBG ", log.Ldate|log.Ltime)
		Info = log.New(ioutil.Discard, "INF ", log.Ldate|log.Ltime)

	case "error":
		Trace = log.New(ioutil.Discard, "TRC ", log.Ldate|log.Ltime)
		Debug = log.New(ioutil.Discard, "DBG ", log.Ldate|log.Ltime)
		Info = log.New(ioutil.Discard, "INF ", log.Ldate|log.Ltime)
		Warning = log.New(ioutil.Discard, "WRN ", log.Ldate|log.Ltime)

	default:
		return errors.New(fmt.Sprintf("Invalid log level: %s", level))

	}

	return nil
}

func Traceln(v ...interface{}) {
	Trace.Println(v...)
}
func Tracef(format string, v ...interface{}) {
	Trace.Printf(format, v...)
}

func Debugln(v ...interface{}) {
	Debug.Println(v...)
}
func Debugf(format string, v ...interface{}) {
	Debug.Printf(format, v...)
}

func Infoln(v ...interface{}) {
	Info.Println(v...)
}
func Infof(format string, v ...interface{}) {
	Info.Printf(format, v...)
}

func Warningln(v ...interface{}) {
	Warning.Println(v...)
}
func Warningf(format string, v ...interface{}) {
	Warning.Printf(format, v...)
}

func Errorln(v ...interface{}) {
	Error.Println(v...)
}
func Errorf(format string, v ...interface{}) {
	Error.Printf(format, v...)
}

func Fatalln(v ...interface{}) {
	Fatal.Println(v...)
}
func Fatalf(format string, v ...interface{}) {
	Fatal.Printf(format, v...)
}

// For backward compatibility
type Logger struct {
	Trace   *log.Logger
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
}

func NewLogger(
	traceHandle io.Writer,
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) *Logger {

	return &Logger{
		Trace:   log.New(traceHandle, "TRC ", log.Ldate|log.Ltime),
		Debug:   log.New(traceHandle, "DBG ", log.Ldate|log.Ltime),
		Info:    log.New(infoHandle, "INF ", log.Ldate|log.Ltime),
		Warning: log.New(warningHandle, "WRN ", log.Ldate|log.Ltime),
		Error:   log.New(errorHandle, "ERR ", log.Ldate|log.Ltime),
		Fatal:   log.New(errorHandle, "FTL ", log.Ldate|log.Ltime),
	}
}

func NewStdLogger() *Logger {
	return NewLogger(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stderr)
}

func (l *Logger) SetLogLevel(level string) error {
	switch level {
	case "trace":
		break
	case "debug":
		l.Trace = log.New(ioutil.Discard, "TRC ", log.Ldate|log.Ltime)
		break
	case "info":
		l.Trace = log.New(ioutil.Discard, "TRC ", log.Ldate|log.Ltime)
		l.Debug = log.New(ioutil.Discard, "DBG ", log.Ldate|log.Ltime)
		break
	case "warning":
		l.Trace = log.New(ioutil.Discard, "TRC ", log.Ldate|log.Ltime)
		l.Debug = log.New(ioutil.Discard, "DBG ", log.Ldate|log.Ltime)
		l.Info = log.New(ioutil.Discard, "INF ", log.Ldate|log.Ltime)
		break
	case "error":
		l.Trace = log.New(ioutil.Discard, "TRC ", log.Ldate|log.Ltime)
		l.Debug = log.New(ioutil.Discard, "DBG ", log.Ldate|log.Ltime)
		l.Info = log.New(ioutil.Discard, "INF ", log.Ldate|log.Ltime)
		l.Warning = log.New(ioutil.Discard, "WRN ", log.Ldate|log.Ltime)
		break
	default:
		return errors.New(fmt.Sprintf("Invalid log level: %s", level))
	}
	return nil
}

func (l *Logger) Traceln(v ...interface{}) {
	l.Trace.Println(v...)
}
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Trace.Printf(format, v...)
}

func (l *Logger) Debugln(v ...interface{}) {
	l.Debug.Println(v...)
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Debug.Printf(format, v...)
}

func (l *Logger) Infoln(v ...interface{}) {
	l.Info.Println(v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Info.Printf(format, v...)
}

func (l *Logger) Warningln(v ...interface{}) {
	l.Warning.Println(v...)
}
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Warning.Printf(format, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.Error.Println(v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Error.Printf(format, v...)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.Fatal.Println(v...)
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Fatal.Printf(format, v...)
}
