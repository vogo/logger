// Copyright 2019 The vogo Authors. All rights reserved.
// author: wongoo

package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
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

func TestSetWriter(t *testing.T) {
	logFile := os.TempDir() + "/test_golang_logger.WriteLog"
	defer os.Remove(logFile)

	f, _ := os.Create(logFile)

	SetOutput(f)

	Info("hello")

	data, _ := os.ReadFile(logFile)
	if !bytes.HasSuffix(data, []byte("hello\n")) {
		t.Errorf("unexpect WriteLog data: %s", data)
	}
}

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("20060102 15:04:05.999"))
	fmt.Println(now.Format("20060102 15:04:05.999999"))
}

func BenchmarkInfo(b *testing.B) {
	SetOutput(io.Discard)
	SetFlags(Lnone)
	for i := 0; i < b.N; i++ {
		Info("hello world")
	}
}

func BenchmarkInfoParallel(b *testing.B) {
	SetOutput(io.Discard)
	SetFlags(Lnone)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("hello world")
		}
	})
}

func BenchmarkInfoWithCaller(b *testing.B) {
	SetOutput(io.Discard)
	SetFlags(Lfile)
	for i := 0; i < b.N; i++ {
		Info("hello world")
	}
}

func BenchmarkInfof(b *testing.B) {
	SetOutput(io.Discard)
	SetFlags(Lnone)
	for i := 0; i < b.N; i++ {
		Infof("%s %s", "hello", "world")
	}
}

func BenchmarkInfofParallel(b *testing.B) {
	SetOutput(io.Discard)
	SetFlags(Lnone)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Infof("%s %s", "hello", "world")
		}
	})
}

func BenchmarkInfofWithCaller(b *testing.B) {
	SetOutput(io.Discard)
	SetFlags(Lfile)
	for i := 0; i < b.N; i++ {
		Infof("%s %s", "hello", "world")
	}
}

func BenchmarkInfofWithCallerParallel(b *testing.B) {
	SetOutput(io.Discard)
	SetFlags(Lfile)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Infof("%s %s", "hello", "world")
		}
	})
}
