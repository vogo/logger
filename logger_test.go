// Copyright 2019 The vogo Authors. All rights reserved.
// author: wongoo

package logger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func TestLogger(t *testing.T) {
	SetFlags(LfileFunc)

	Trace("trace", "trace")
	Debug("debug", "debug")
	Info("info", "info")
	Warn("warn", "warn")
	Error("error", "error")

	Tracef("%s-%s", "trace", "trace")
	Debugf("%s-%s", "debug", "debug")
	Infof("%s-%s", "info", "info")
	Warnf("%s-%s", "warn", "warn")
	Errorf("%s-%s", "error", "error")
}

func TestLog(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("hello world")
}

func TestSetWriter(t *testing.T) {
	logFile := "/tmp/test_golang_logger.writeLog"
	defer os.Remove(logFile)

	SetOutput(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // megabytes after which new file is created
		MaxBackups: 10, // number of backups
		MaxAge:     30, // days
	})

	Info("hello")

	data, _ := ioutil.ReadFile(logFile)
	if !bytes.HasSuffix(data, []byte("hello\n")) {
		t.Errorf("unexpect writeLog data: %s", data)
	}
}

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("20060102 15:04:05.999"))
	fmt.Println(now.Format("20060102 15:04:05.999999"))
}

func BenchmarkInfo(b *testing.B) {
	SetOutput(ioutil.Discard)
	SetFlags(Lnone)
	for i := 0; i < b.N; i++ {
		Info("hello world")
	}
}

func BenchmarkInfoWithCaller(b *testing.B) {
	SetOutput(ioutil.Discard)
	SetFlags(Lfile)
	for i := 0; i < b.N; i++ {
		Info("hello world")
	}
}

func BenchmarkInfof(b *testing.B) {
	SetOutput(ioutil.Discard)
	SetFlags(Lnone)
	for i := 0; i < b.N; i++ {
		Infof("%s %s", "hello", "world")
	}
}

func BenchmarkInfofWithCaller(b *testing.B) {
	SetOutput(ioutil.Discard)
	SetFlags(Lfile)
	for i := 0; i < b.N; i++ {
		Infof("%s %s", "hello", "world")
	}
}

// -------------- internal log ----------------------

func BenchmarkLogPrintln(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	for i := 0; i < b.N; i++ {
		log.Println("hello world")
	}
}

func BenchmarkLogPrintlnCaller(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	for i := 0; i < b.N; i++ {
		log.Println("hello world")
	}
}

func BenchmarkLogPrintf(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	for i := 0; i < b.N; i++ {
		log.Printf("%s %s", "hello", "world")
	}
}

func BenchmarkLogPrintfCaller(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	for i := 0; i < b.N; i++ {
		log.Printf("%s %s", "hello", "world")
	}
}
