package log

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
)

type FileEncoder struct {
}

func NewFileCore(name string, writers ...io.Writer) *Core {
	if name == "" {
		name = "app"
	}
	_ = os.MkdirAll("./logs", 0755)
	allFileName := fmt.Sprintf("logs/%s.log", name)
	errFileName := fmt.Sprintf("logs/%s_err.log", name)
	core := &Core{
		encoder: new(FileEncoder),
	}
	if len(writers) > 0 {
		core.infoWriter = writers[0]
		if len(writers) > 1 {
			core.errWriter = writers[1]
		} else {
			core.errWriter = writers[0]
		}
	} else {
		core.allWriter = &lumberjack.Logger{
			Filename:   allFileName, // 日志文件路径
			MaxSize:    1,           // 日志文件最大大小(MB)
			MaxBackups: 100,         // 保留旧文件的最大个数
			MaxAge:     28,          // 保留旧文件的最大天数
			Compress:   false,       // 是否压缩/归档旧文件
		}
		core.errWriter = &lumberjack.Logger{
			Filename:   errFileName, // 日志文件路径
			MaxSize:    1,           // 日志文件最大大小(MB)
			MaxBackups: 10,          // 保留旧文件的最大个数
			MaxAge:     28,          // 保留旧文件的最大天数
			Compress:   false,       // 是否压缩/归档旧文件
		}
	}

	return core
}

func (e *FileEncoder) Encode(ac *Entry, params []any) string {
	if ac == nil {
		return ""
	}

	w := bufferPool.Get().(*Buffer)
	defer w.close()

	if ac.Date != "" {
		w.WriteString(ac.Date + " ")
	}
	w.WriteString(ac.Level.String())
	if len(ac.Fields) > 0 {
		w.WriteString(" [" + strings.Join(ac.Fields, " ") + "]")
	}

	if params != nil {
		w.WriteString(" ")
		format, ok := params[0].(string)
		if ok && strings.ContainsRune(format, '%') {
			if _, err := fmt.Fprintf(w, format, params[1:]...); err != nil {
				return ""
			}
		} else {
			if _, err := fmt.Fprint(w, params...); err != nil {
				return ""
			}
		}
	}
	if ac.AddCaller && ac.Caller != "" {
		w.WriteString(" " + ac.Caller)
	}

	for i, layer := range ac.Stack {
		if _, err := fmt.Fprintf(w, "\n %2d %v", i+1, layer); err != nil {
			return ""
		}
	}

	return w.String()
}
