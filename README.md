# log

[//]: # (<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">)

[![Build Status](https://github.com/yulefox/log/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/yulefox/log/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/yulefox/log/branch/master/graph/badge.svg)](https://codecov.io/gh/yulefox/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/yulefox/log)](https://goreportcard.com/report/github.com/yulefox/log)
[![GoDoc](https://pkg.go.dev/badge/github.com/yulefox/log?status.svg)](https://pkg.go.dev/github.com/yulefox/log?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/yulefox/log/-/badge.svg)](https://sourcegraph.com/github.com/yulefox/log?badge)
[![Open Source Helpers](https://www.codetriage.com/yulefox/log/badges/users.svg)](https://www.codetriage.com/yulefox/log)
[![Release](https://img.shields.io/github/release/yulefox/log.svg?style=flat-square)](https://github.com/yulefox/log/releases)

[//]: # ([![TODOs]&#40;https://badgen.net/https/api.tickgit.com/badgen/github.com/yulefox/log&#41;]&#40;https://www.tickgit.com/browse?repo=github.com/yulefox/log&#41;)


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

## License

MIT
