package log

import (
	"fmt"
	"io"
	"os"
	"time"

	. "github.com/lollipopkit/gommon/res"
)

const (
	warn    = YELLOW + "[WAR] " + NOCOLOR
	err     = RED + "[ERR] " + NOCOLOR
	info    = CYAN + "[INF] " + NOCOLOR
	success = GREEN + "[SUC] " + NOCOLOR
)

type Config struct {
	// Debug is the flag to enable debug log.
	Debug bool
	// PrintTime is the flag to print time.
	PrintTime bool
	// LogDir is the directory to store log files.
	// If you set this to empty string (default), log files will not be created.
	LogDir string
	// FilePerm is the permission of log files.
	FilePerm os.FileMode
}

var (
	config Config
	writer io.Writer
	ticker *time.Ticker
)

func Setup(config_ Config) {
	config = config_
	if config.FilePerm == 0 {
		config.FilePerm = 0755
	}

	if err := os.MkdirAll(config.LogDir, config.FilePerm); err != nil {
		panic(err)
	}

	go setup()
}

func printf(format string, args ...any) {
	if config.PrintTime {
		format = time.Now().Format("2006-01-02 15:04:05") + " " + format
	}
	f := fmt.Sprintf(format+"\n", args...)
	if writer == nil {
		print(f)
	} else {
		writer.Write([]byte(f))
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
	if !config.Debug {
		return
	}
	printf("[DEBUG] "+format, args...)
}

// Must call this func using:
// `go setup()`
func setup() {
	if ticker != nil {
		ticker.Stop()
		ticker = nil
	}
	// `logDir` not set
	// no need to set `writer`
	if len(config.LogDir) == 0 {
		return
	}
	if ticker == nil {
		ticker = time.NewTicker(time.Minute)
	}

	for range ticker.C {
		file := config.LogDir + time.Now().Format("2006-01-02") + ".txt"
		logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, config.FilePerm)
		if err != nil {
			panic(err)
		}
		writer = io.MultiWriter(os.Stdout, logFile)
	}
}
