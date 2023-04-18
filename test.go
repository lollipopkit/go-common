package main

import (
	"time"

	"github.com/lollipopkit/gommon/term"
)

func main() {
	spinner := term.NewSpinner()
	spinner.SetString("Loading...\n")
	time.Sleep(3 * time.Second)
	spinner.SetString("Fetching...")
	time.Sleep(3 * time.Second)
	spinner.Stop(false)
}
