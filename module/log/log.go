package log

import (
	"github.com/charmbracelet/log"
	"os"
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
	Logger.SetLevel(log.DebugLevel)
}
