package term

import (
	"errors"
	"strings"
	"time"

	"atomicgo.dev/cursor"
)

var (
	Frames1 = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	Frames2 = []string{"-", "\\", "|", "/"}
	Frames3 = []string{"◜", "◠", "◝", "◞", "◡", "◟"}

	ErrSpinnerAlreadyStarted = errors.New("spinner already started")
)

// spinner is a simple spinner that can be used to indicate progress.
type spinner struct {
	// Frames is the list of frames to use for the spinner.
	Frames []string
	// Index is the current index of the spinner.
	Index int
	// Interval is the interval to use for the spinner.
	Interval time.Duration
	// Suffix is the suffix to use for the spinner.
	Suffix string
	// Ticker is the ticker used for the spinner.
	Ticker *time.Ticker
}

// NewSpinner returns a new spinner.
func NewSpinner(frames []string, interval time.Duration) *spinner {
	return &spinner{
		Frames: frames,
		Interval: interval,
	}
}

// Stop stops the spinner.
func (s *spinner) Stop() {
	s.Ticker.Stop()
}

// Start starts the spinner.
func (s *spinner) Start() error {
	if s.Ticker != nil {
		return ErrSpinnerAlreadyStarted
	}

	s.Ticker = time.NewTicker(s.Interval)
	go func() {
		for range s.Ticker.C {
			s.Index = (s.Index + 1) % len(s.Frames)
			cursor.StartOfLine()
			print(s.Frames[s.Index]  + " " + strings.TrimSpace(s.Suffix))
		}
	}()

	return nil
}

// SetString sets the suffix of the spinner.
func (s *spinner) SetString(suffix string) {
	s.Suffix = suffix
}