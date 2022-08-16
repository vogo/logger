// Copyright 2019 The vogo Authors. All rights reserved.
// author: wongoo

package logger

import (
	"fmt"
	"io"
	"log"
	"testing"
)

// StdLogPrintf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
// Not check whether the output writer is discard.
func StdLogPrintf(format string, v ...any) {
	_ = log.Output(2, fmt.Sprintf(format, v...))
}

// StdLogPrintln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
// Not check whether the output writer is discard.
func StdLogPrintln(v ...any) {
	_ = log.Output(2, fmt.Sprintln(v...))
}

func TestLog(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("hello world")
}

func BenchmarkLogPrintln(b *testing.B) {
	log.SetOutput(io.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	for i := 0; i < b.N; i++ {
		StdLogPrintln("hello world")
	}
}

func BenchmarkLogPrintlnParallel(b *testing.B) {
	log.SetOutput(io.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			StdLogPrintln("hello world")
		}
	})
}

func BenchmarkLogPrintlnCaller(b *testing.B) {
	log.SetOutput(io.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	for i := 0; i < b.N; i++ {
		StdLogPrintln("hello world")
	}
}

func BenchmarkLogPrintf(b *testing.B) {
	log.SetOutput(io.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	for i := 0; i < b.N; i++ {
		StdLogPrintf("%s %s", "hello", "world")
	}
}

func BenchmarkLogPrintfCaller(b *testing.B) {
	log.SetOutput(io.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	for i := 0; i < b.N; i++ {
		StdLogPrintf("%s %s", "hello", "world")
	}
}

func BenchmarkLogPrintfCallerParallel(b *testing.B) {
	log.SetOutput(io.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			StdLogPrintf("%s %s", "hello", "world")
		}
	})
}
