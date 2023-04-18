package term

import (
	"testing"
	"time"
)

func TestSpinner(t *testing.T) {
	spinner := NewSpinner()
	spinner.SetString("Loading...\n")
	time.Sleep(3 * time.Second)
	spinner.SetString("Fetching...")
	time.Sleep(3 * time.Second)
	spinner.Stop()
}
