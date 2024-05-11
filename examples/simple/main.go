package main

import (
	"fmt"
	"sync"

	"github.com/yulefox/log"
)

func main() {
	var wg sync.WaitGroup
	log.Info("tag", "Here is a simple example.")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			tag := fmt.Sprintf("foo-%d", i)
			for j := 0; j < 300; j++ {
				log.Debug(tag, "This is a debug log entry ", i, i+1)
				log.Info(tag, "This is an info log entry ", i)
				log.Warn(tag, "This is a warning log entry")
				log.Error(tag, "This is an error log entry with caller stack")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	defer func() {
		if err := recover(); err != nil {
			log.Warn("recover", "Recovered from panic: ", err)
		}
	}()
	//log.Panic("panic", "This is a panic with caller stack.")
	log.Fatal("fatal", "This should not be logged.")
	log.Info("info", "This should not be logged.")
}
