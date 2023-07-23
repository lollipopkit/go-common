package log_test

import (
	"testing"

	"github.com/lollipopkit/gommon/log"
)

func TestLog(t *testing.T) {
	log.Debug("debug")
	log.Setup(log.Config{
		PrintTime: true,
		LogPath:   "log/",
		Debug:     true,
	})
	log.Info("info")
	log.Suc("suc")
	log.Warn("warn")
	log.Err("err")
	log.Debug("debug")
}
