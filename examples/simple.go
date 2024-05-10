package main

import "github.com/yulefox/log"

func main() {
	// default logger should be invoked with a tag
	log.Info("tag", "Here is a simple example.")
	log.Debug("main", "This is a debug message.")
	log.Info("example", "Hello, %s!", "world")
	log.Error("debug", "This is an error with caller stack.")
}
