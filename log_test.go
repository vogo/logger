// Copyright 2019 The vogo Authors. All rights reserved.
// author: wongoo

package logger

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("hello world")
}

func BenchmarkLogPrintln(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	for i := 0; i < b.N; i++ {
		log.Println("hello world")
	}
}

func BenchmarkLogPrintlnParallel(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Println("hello world")
		}
	})
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

func BenchmarkLogPrintfCallerParallel(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Printf("%s %s", "hello", "world")
		}
	})
}
