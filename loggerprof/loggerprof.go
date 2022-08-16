package main

import (
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/vogo/logger"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	logger.SetOutput(io.Discard)
	logger.SetFlags(logger.Lfile)

	for n := 0; n < 1024; n++ {
		go func() {
			for i := 0; i < 100000; i++ {
				logger.WriteLog("a", "b")
			}
		}()
	}

	select {}
}
