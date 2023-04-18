package main

import (
	"time"

	"github.com/lollipopkit/gommon/term"
)

func main() {
	spinner := term.NewSpinner()
	spinner.SetString("Loading...\n")
	spinner.Start()
	time.Sleep(3 * time.Second)
	spinner.Stop()
}
