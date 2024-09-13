# log

[//]: # (<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">)

[![Build Status](https://github.com/yulefox/log/workflows/GoTest/badge.svg)](https://github.com/yulefox/log/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/yulefox/log/branch/main/graph/badge.svg)](https://codecov.io/gh/yulefox/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/yulefox/log)](https://goreportcard.com/report/github.com/yulefox/log)

[//]: # ([![Sourcegraph]&#40;https://sourcegraph.com/github.com/yulefox/log/-/badge.svg&#41;]&#40;https://sourcegraph.com/github.com/yulefox/log?badge&#41;)
[//]: # ([![Open Source Helpers]&#40;https://www.codetriage.com/yulefox/log/badges/users.svg&#41;]&#40;https://www.codetriage.com/yulefox/log&#41;)

[![GoDoc](https://pkg.go.dev/badge/github.com/yulefox/log?status.svg)](https://pkg.go.dev/github.com/yulefox/log?tab=doc)
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
	log.Info("Here is a simple example.")
	log.Debug("This is a debug message.")
	log.Info("Hello, %s!", "world")
	log.Error("This is an error with caller stack.")
	log.Panic("This is a panic with caller stack.")
	//log.Fatal("This should not be logged.")
	//log.Info("This should not be logged.")
}
```

[//]: # (## Features)

[//]: # ()
[//]: # ()
[//]: # (UUIDs are 16 bytes &#40;128 bits&#41; and 36 chars as string representation. Twitter Snowflake)

[//]: # (ids are 8 bytes &#40;64 bits&#41; but require machine/data-center configuration and/or central)

[//]: # (generator servers. xid stands in between with 12 bytes &#40;96 bits&#41; and a more compact)

[//]: # (URL-safe string representation &#40;20 chars&#41;. No configuration or central generator server)

[//]: # (is required. So it can be used directly in server's code.)

[//]: # ()
[//]: # (| Name     | Binary Size | String Size    | Features                         |)

[//]: # (|----------|-------------|----------------|----------------------------------|)

[//]: # (| [go log] | 16 bytes    | 36 chars       | configuration free, not sortable |)

[//]: # (| [LogRus] | 16 bytes    | 22 chars       | configuration free, not sortable |)

[//]: # (| log      | 12 bytes    | 20 chars       | configuration free               |)

[//]: # ()
[//]: # ([UUID]: https://en.wikipedia.org/wiki/Universally_unique_identifier)

[//]: # ([shortuuid]: https://github.com/stochastic-technologies/shortuuid)

[//]: # ([Snowflake]: https://blog.twitter.com/2010/announcing-snowflake)

[//]: # ([MongoID]: https://docs.mongodb.org/manual/reference/object-id/)

## License

MIT
