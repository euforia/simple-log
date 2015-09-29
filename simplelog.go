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

type Logger struct {
	Trace   *log.Logger
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func NewLogger(
	traceHandle io.Writer,
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) *Logger {

	return &Logger{
		Trace:   log.New(traceHandle, "TRC: ", log.Ldate|log.Ltime|log.Lshortfile),
		Debug:   log.New(traceHandle, "DBG: ", log.Ldate|log.Ltime|log.Lshortfile),
		Info:    log.New(infoHandle, "INF: ", log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(warningHandle, "WRN: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error:   log.New(errorHandle, "ERR: ", log.Ldate|log.Ltime|log.Lshortfile),
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
		l.Trace = log.New(ioutil.Discard, "TRC: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	case "info":
		l.Trace = log.New(ioutil.Discard, "TRC: ", log.Ldate|log.Ltime|log.Lshortfile)
		l.Debug = log.New(ioutil.Discard, "DBG: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	case "warning":
		l.Trace = log.New(ioutil.Discard, "TRC: ", log.Ldate|log.Ltime|log.Lshortfile)
		l.Debug = log.New(ioutil.Discard, "DBG: ", log.Ldate|log.Ltime|log.Lshortfile)
		l.Info = log.New(ioutil.Discard, "INF: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	case "error":
		l.Trace = log.New(ioutil.Discard, "TRC: ", log.Ldate|log.Ltime|log.Lshortfile)
		l.Debug = log.New(ioutil.Discard, "DBG: ", log.Ldate|log.Ltime|log.Lshortfile)
		l.Info = log.New(ioutil.Discard, "INF: ", log.Ldate|log.Ltime|log.Lshortfile)
		l.Warning = log.New(ioutil.Discard, "WRN: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	default:
		return errors.New(fmt.Sprintf("Invalid log level: %s", level))
	}
	return nil
}
