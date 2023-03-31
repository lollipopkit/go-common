package term

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	_debug, _inited bool
	_logDir         string
	_perm           os.FileMode
)

const (
	RED     = "\033[91m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[93m"
	CYAN    = "\033[96m"
	NOCOLOR = "\033[0m"
)

const (
	warn    = YELLOW + "[WARN] " + NOCOLOR
	err     = RED + "[ERROR] " + NOCOLOR
	info    = CYAN + "[INFO] " + NOCOLOR
	success = GREEN + "[SUCCESS] " + NOCOLOR
)

func Config(logDir string, perm os.FileMode) {
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

func Warn(format string, args ...any) {
	log.Printf(warn+format, args...)
}

func Info(format string, args ...any) {
	log.Printf(info+format, args...)
}

func Err(format string, args ...any) {
	log.Printf(err+format, args...)
}

func Suc(format string, args ...any) {
	log.Printf(success+format, args...)
}

func Debug(format string, args ...any) {
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
