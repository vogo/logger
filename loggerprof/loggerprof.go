package main

import (
	"github.com/vogo/logger"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	logger.SetOutput(ioutil.Discard)
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
