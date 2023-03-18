package logger

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	_debug, _inited bool
	_logDir string
	_perm os.FileMode
)

func SetPath(logDir string, perm os.FileMode) {
	if err := os.MkdirAll(_logDir, _perm); err != nil {
		panic(err)
	}

	_logDir = logDir
	_perm = perm
	_inited = true

	go setup()
}

func SetDebug(debug bool) {
	_debug = debug
}

func W(format string, args ...any) {
	log.Printf("[WARN] "+format, args...)
}

func I(format string, args ...any) {
	log.Printf("[INFO] "+format, args...)
}

func E(format string, args ...any) {
	log.Printf("[ERROR] "+format, args...)
}

func D(format string, args ...any) {
	if !_debug {
		return
	}
	log.Printf("[DEBUG] "+format, args...)
}

// Must call this func using:
// `go setup()`
func setup() {
	for {
		file := _logDir + time.Now().Format("2006-01-02") + ".txt"
		logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, _perm)
		if err != nil {
			panic(err)
		}
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multiWriter)
		time.Sleep(time.Hour)
	}
}