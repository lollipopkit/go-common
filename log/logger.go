package log

import (
	"fmt"
	"io"
	"os"
	"time"

	. "github.com/lollipopkit/gommon/res"
)

const (
	warn    = YELLOW + "[WAR]" + NOCOLOR + " "
	err     = RED + "[ERR]" + NOCOLOR + " "
	info    = CYAN + "[INF]" + NOCOLOR + " "
	success = GREEN + "[SUC]" + NOCOLOR + " "
	debug   = MAGENTA + "[DEBUG]" + NOCOLOR + " "
)

type Config struct {
	// Debug is the flag to enable debug log.
	Debug bool
	// PrintTime is the flag to print time.
	PrintTime bool
	// LogPath is the directory to store log files.
	//
	// If you set this to empty string (default), log files will not be created.
	//
	// If it not ends with "/", "/" will NOT be appended automatically.
	LogPath string
	// LogFileNameFormat is the format of log file name.
	//
	// Default is "2006-01-02".
	LogFileNameFormat string
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

	if len(config.LogPath) == 0 {
		return
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
	printf(YELLOW+format+NOCOLOR, args...)
}

func Info(format string, args ...any) {
	printf(info+format, args...)
}

func Cyan(format string, args ...any) {
	printf(CYAN+format+NOCOLOR, args...)
}

func Err(format string, args ...any) {
	printf(err+format, args...)
}

func Red(format string, args ...any) {
	printf(RED+format+NOCOLOR, args...)
}

func Suc(format string, args ...any) {
	printf(success+format, args...)
}

func Green(format string, args ...any) {
	printf(GREEN+format+NOCOLOR, args...)
}

func Debug(format string, args ...any) {
	if !config.Debug {
		return
	}
	printf(debug+format, args...)
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
	if len(config.LogPath) == 0 {
		return
	}
	if ticker == nil {
		ticker = time.NewTicker(time.Minute)
	}

	if err := os.MkdirAll(config.LogPath, config.FilePerm); err != nil {
		panic(err)
	}

	// Because ticker will not tick immediately,
	// we need to set writer manually once before ticker starts.
	setWritter()
	for range ticker.C {
		setWritter()
	}
}

func setWritter() {
	timeFmt := "2006-01-02"
	if len(config.LogFileNameFormat) > 0 {
		timeFmt = config.LogFileNameFormat
	}
	file := config.LogPath + time.Now().Format(timeFmt) + ".txt"
	logFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, config.FilePerm)
	if err != nil {
		panic(err)
	}
	writer = io.MultiWriter(os.Stdout, logFile)
}
