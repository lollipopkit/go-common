package main

import (
	"testing"
	"time"

	"github.com/lollipopkit/gommon/term"
)

func TestTermSpinner(t *testing.T) {
	spinner := term.NewSpinner()
	spinner.SetString("Loading...\n")
	spinner.Start(77*time.Millisecond)
	time.Sleep(3 * time.Second)
	spinner.Stop()
}
