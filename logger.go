// author: wongoo

package logger

import (
	"fmt"
	"time"
)

//Logger common log interface
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})

	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})

	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

// DefaultLogger uses the stdlib log package for logging.
var DefaultLogger Logger = defaultLogger{}

type defaultLogger struct{}

func p(level, output string) {
	now := time.Now().Format("20060102 15:04:05.99999")
	fmt.Printf("%-25s [%-5s] %s\n", now, level, output)
}

func (defaultLogger) Info(args ...interface{}) {
	p("INFO", fmt.Sprint(args...))
}
func (defaultLogger) Warn(args ...interface{}) {
	p("WARN", fmt.Sprint(args...))
}
func (defaultLogger) Error(args ...interface{}) {
	p("ERROR", fmt.Sprint(args...))
}
func (defaultLogger) Debug(args ...interface{}) {
	p("DEBUG", fmt.Sprint(args...))
}

func (defaultLogger) Infof(format string, args ...interface{}) {
	p("INFO", fmt.Sprintf(format, args...))
}
func (defaultLogger) Warnf(format string, args ...interface{}) {
	p("WARN", fmt.Sprintf(format, args...))
}
func (defaultLogger) Errorf(format string, args ...interface{}) {
	p("ERROR", fmt.Sprintf(format, args...))
}
func (defaultLogger) Debugf(format string, args ...interface{}) {
	p("DEBUG", fmt.Sprintf(format, args...))
}

func (defaultLogger) Printf(format string, args ...interface{}) {
	p("INFO", fmt.Sprintf(format, args...))
}

func (defaultLogger) Println(args ...interface{}) {
	p("INFO", fmt.Sprint(args...))
}
