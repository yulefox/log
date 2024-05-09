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
	// 使用示例
	for i := 0; i < 100; i++ {
		go func() {
			log.Debug("debug")
			log.Info("info")
			log.Warn("warn")
			log.Error("error")
			log.Fatal("fatal")
		}
    }
}
```

## Docs
    
    [![GoDoc](https://godoc.org/github.com/yulefox/log?status.svg)](https://godoc.org/github.com/yulefox/log?status.svg)

## License

MIT

## Contributing

## Changelog

## Credits

## Contact

## References

## See Also

## Acknowledgements

## Appendix

## FAQ

## License

## Support

## Troubleshooting

## Versioning

## Authors

## Acknowledgments

