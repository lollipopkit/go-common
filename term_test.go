package main

import (
	"testing"
	"time"

	"github.com/lollipopkit/gommon/term"
)

func TestTermSpinner(t *testing.T) {
	spinner := term.NewSpinner("Loading...\n", term.Frames1, 77*time.Millisecond)
	spinner.Start()
	time.Sleep(3 * time.Second)
	spinner.Stop()
}