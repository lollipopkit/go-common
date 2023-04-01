package term

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	_debug, _printTime bool
	_logDir            string
	_perm              os.FileMode
	_writer            io.Writer
	_ticker            *time.Ticker
	_interval          = time.Minute
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

func Config(logDir string, perm os.FileMode, printTime bool) {
	if err := os.MkdirAll(_logDir, _perm); err != nil {
		panic(err)
	}

	_logDir = logDir
	_perm = perm
	_printTime = printTime

	go setup()
}

func SetDebug(debug bool) {
	_debug = debug
}

func printf(format string, args ...any) {
	if _printTime {
		format = time.Now().Format("2006-01-02 15:04:05") + " " + format
	}
	f := fmt.Sprintf(format+"\n", args...)
	if _writer == nil {
		print(f)
	} else {
		_writer.Write([]byte(f))
	}
}

func Warn(format string, args ...any) {
	printf(warn+format, args...)
}

func Yellow(format string, args ...any) {
	fmt.Printf(YELLOW+format+NOCOLOR, args...)
}

func Info(format string, args ...any) {
	printf(info+format, args...)
}

func Cyan(format string, args ...any) {
	fmt.Printf(CYAN+format+NOCOLOR, args...)
}

func Err(format string, args ...any) {
	printf(err+format, args...)
}

func Red(format string, args ...any) {
	fmt.Printf(RED+format+NOCOLOR, args...)
}

func Suc(format string, args ...any) {
	printf(success+format, args...)
}

func Green(format string, args ...any) {
	fmt.Printf(GREEN+format+NOCOLOR, args...)
}

func Debug(format string, args ...any) {
	if !_debug {
		return
	}
	printf("[DEBUG] "+format, args...)
}

// Must call this func using:
// `go setup()`
func setup() {
	if _ticker == nil {
		_ticker = time.NewTicker(_interval)
	} else {
		_ticker.Reset(_interval)
	}

	for range _ticker.C {
		file := _logDir + time.Now().Format("2006-01-02") + ".txt"
		logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, _perm)
		if err != nil {
			panic(err)
		}
		_writer = io.MultiWriter(os.Stdout, logFile)
	}
}
