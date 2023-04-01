package main

import (
	"time"

	"github.com/lollipopkit/gommon/term"
)

func main() {
	spinner := term.NewSpinner()
	spinner.SetString("Loading...\n")
	spinner.Start(77*time.Millisecond)
	time.Sleep(2 * time.Second)
	spinner.Stop()
}
