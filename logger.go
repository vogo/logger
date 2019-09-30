// Copyright 2019 The vogo Authors. All rights reserved.
// author: wongoo

package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	TagTrace = "TRAC"
	TagDebug = "DEBG"
	TagInfo  = "INFO"
	TagWarn  = "WARN"
	TagError = "ERRO"
	TagFatal = "FATL"
	TagPanic = "PNIC"
	TagPrint = "PRNT"

	LevelTrace = 5
	LevelDebug = 4
	LevelInfo  = 3
	LevelWarn  = 2
	LevelError = 1
	LevelFatal = 0
)

const (
	Lnone     = 0             // none file
	Lfile     = 1             // d.go:23
	Lfunc     = 1 << 1        // foo
	LfileFunc = Lfunc | Lfile // d.go:foo:23
)

var (
	level            = LevelInfo
	output io.Writer = os.Stdout
	lock   sync.Mutex
	buf    []byte
	flag   int
)

// SetLevel set logger level
func SetLevel(l int) {
	level = l
}

// SetOutput set logger output writer
func SetOutput(w io.Writer) {
	output = w
}

// SetFlags set logger flags
func SetFlags(f int) {
	flag = f
}

// Writer return the logger writer
func Writer() io.Writer {
	return output
}

func Trace(a ...interface{}) {
	if level < LevelTrace {
		return
	}
	WriteLog(TagTrace, fmt.Sprint(a...))
}

func Debug(a ...interface{}) {
	if level < LevelDebug {
		return
	}
	WriteLog(TagDebug, fmt.Sprint(a...))
}

func Info(a ...interface{}) {
	if level < LevelInfo {
		return
	}
	WriteLog(TagInfo, fmt.Sprint(a...))
}

func Warn(a ...interface{}) {
	if level < LevelWarn {
		return
	}
	WriteLog(TagWarn, fmt.Sprint(a...))
}

func Error(a ...interface{}) {
	if level < LevelError {
		return
	}
	WriteLog(TagError, fmt.Sprint(a...))
}

func Tracef(format string, a ...interface{}) {
	if level < LevelTrace {
		return
	}
	WriteLog(TagTrace, fmt.Sprintf(format, a...))
}

func Debugf(format string, a ...interface{}) {
	if level < LevelDebug {
		return
	}
	WriteLog(TagDebug, fmt.Sprintf(format, a...))
}

func Infof(format string, a ...interface{}) {
	if level < LevelInfo {
		return
	}
	WriteLog(TagInfo, fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...interface{}) {
	if level < LevelWarn {
		return
	}
	WriteLog(TagWarn, fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...interface{}) {
	if level < LevelError {
		return
	}
	WriteLog(TagError, fmt.Sprintf(format, a...))
}

func Fatal(a ...interface{}) {
	WriteLog(TagFatal, fmt.Sprint(a...))
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	WriteLog(TagFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func Fatalln(a ...interface{}) {
	WriteLog(TagFatal, fmt.Sprint(a...))
	os.Exit(1)
}

func Print(a ...interface{}) {
	WriteLog(TagPrint, fmt.Sprint(a...))
}

func Printf(format string, a ...interface{}) {
	WriteLog(TagPrint, fmt.Sprintf(format, a...))
}

func Println(format string, a ...interface{}) {
	WriteLog(TagPrint, fmt.Sprintf(format, a...))
}

func Panic(a ...interface{}) {
	s := fmt.Sprint(a...)
	WriteLog(TagPanic, s)
	panic(s)
}

func Panicf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	WriteLog(TagPanic, s)
	panic(s)
}

func Panicln(a ...interface{}) {
	s := fmt.Sprint(a...)
	WriteLog(TagPanic, s)
	panic(s)
}

// WriteLog write log data
func WriteLog(tag, s string) {
	var (
		pc       uintptr
		fileName string
		funcName string
		line     int
		callerOk bool
	)

	t := time.Now()

	if flag&Lfile != 0 {
		pc, fileName, line, callerOk = runtime.Caller(2)
		if callerOk {
			for i := len(fileName) - 1; i > 0; i-- {
				if fileName[i] == '/' {
					fileName = fileName[i+1:]
					break
				}
			}
			if flag&Lfunc != 0 {
				funcName = runtime.FuncForPC(pc).Name() // main.(*MyStruct).foo

				for i := len(funcName) - 1; i > 0; i-- {
					if funcName[i] == '.' {
						funcName = funcName[i+1:]
						break
					}
				}
			}
		}
	}

	lock.Lock()
	buf = buf[:0]

	year, month, day := t.Date()
	appendNumber(&buf, year, 4)
	buf = append(buf, '/')
	appendNumber(&buf, int(month), 2)
	buf = append(buf, '/')
	appendNumber(&buf, day, 2)
	buf = append(buf, ' ')

	hour, min, sec := t.Clock()
	appendNumber(&buf, hour, 2)
	buf = append(buf, ':')
	appendNumber(&buf, min, 2)
	buf = append(buf, ':')
	appendNumber(&buf, sec, 2)
	buf = append(buf, '.')
	appendNumber(&buf, t.Nanosecond()/1e6, 3)

	buf = append(buf, ' ')
	buf = append(buf, tag...)

	if flag&Lfile != 0 {
		buf = append(buf, ' ', '[')
		if callerOk {
			buf = append(buf, fileName...)
			if flag&Lfunc != 0 {
				buf = append(buf, ':')
				buf = append(buf, funcName...)
			}
			buf = append(buf, ':')
			appendNumber(&buf, line, -1)
		} else {
			buf = append(buf, '?')
		}

		buf = append(buf, ']')
	}

	buf = append(buf, ' ')
	buf = append(buf, s...)
	if s == "" || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}

	_, _ = output.Write(buf)

	lock.Unlock()
}

// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
func appendNumber(buf *[]byte, i, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}
