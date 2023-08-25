package log

import (
	"fmt"
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/charmbracelet/log"
	"os"
	"runtime"
	"strconv"
	"time"
)

var Logger *log.Logger

func Init() {
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339Nano,
		Prefix:          "Tiktok",
	})

	level := config.GetString("log.level")
	switch level {
	case "debug":
		Logger.SetLevel(log.DebugLevel)
	case "info":
		Logger.SetLevel(log.InfoLevel)
	case "warn":
		Logger.SetLevel(log.WarnLevel)
	case "error":
		Logger.SetLevel(log.ErrorLevel)
	default:
		Logger.SetLevel(log.DebugLevel)
	}
	Logger.Debugf("log level: %v", level)
}

// HandleError 封装了错误打印，高亮错误信息，单元测试时使用！！！
// 相当于 if err != nil { Logger.Error(err) }
func HandleError(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		logMessage := fmt.Sprintf("<%s:%s> %s", file, util.HighlightString(util.RED, strconv.Itoa(line)), err.Error())
		Logger.Errorf(logMessage)
	}
}
