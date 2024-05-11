# log
Another logger for Go.

## Install

```bash
go get github.com/yulefox/log
```

## Usage

```go
package main

import "github.com/yulefox/log"

func main() {
	log.Info("tag", "Here is a simple example.")
	log.Debug("main", "This is a debug message.")
	log.Info("", "Hello, %s!", "world")
	log.Error("", "This is an error with caller stack.")
	log.Panic("panic", "This is a panic with caller stack.")
	//log.Fatal("fatal", "This should not be logged.")
	//log.Info("info", "This should not be logged.")
}
```

## Docs
    
    [![GoDoc](https://godoc.org/github.com/yulefox/log?status.svg)](https://godoc.org/github.com/yulefox/log?status.svg)

## License

MIT
