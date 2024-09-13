package main

import (
	"sync"
	"time"

	"github.com/yulefox/log"
)

func main() {
	log.Init(
		log.SetCaller(false),
		log.SetLevel(log.INFO),
		log.SetTimeFormat("2006-01-02 15:04:05", time.Now().UTC),
		log.AddFileLogger("app"),
		log.AddFileLogger("test"),
		//log.AddJsonLogger(&lumberjack.Logger{
		//	Filename:   "logs/app_json.log", // 日志文件路径
		//	MaxSize:    1,           // 日志文件最大大小(MB)
		//	MaxBackups: 100,         // 保留旧文件的最大个数
		//	MaxAge:     28,          // 保留旧文件的最大天数
		//	Compress:   false,       // 是否压缩/归档旧文件
		//}),
	)

	log.Info("", "Here is a simple example.")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			//tag = ""
			for j := 0; j < 3; j++ {
				log.Debug("This is a debug log entry ", i, i+1)
				log.Info("This is an info log entry ", i)
				log.Warn("This is a warning log entry")
				log.Error("This is an error log entry with caller stack")
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
	log.Panic("This is a panic with caller stack.")
	log.Fatal("This should not be logged.")
	log.Info("This should not be logged.")
}
