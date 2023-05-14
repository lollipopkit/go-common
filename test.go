package main

import (
	"time"

	"github.com/lollipopkit/gommon/term"
)

func main() {
	spinner := term.NewCustomSpinner(term.Frames2, time.Millisecond*77)
	spinner.SetString("Loading...\n")
	time.Sleep(3 * time.Second)
	spinner.SetString("Fetching...")
	time.Sleep(3 * time.Second)
	spinner.Stop(false)
	println("Done!")
}
