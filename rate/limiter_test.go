package rate_test

import (
	"testing"
	"time"

	"github.com/lollipopkit/gommon/rate"
)

func TestRateLimiter(t *testing.T) {
	l := rate.NewLimiter[string](1*time.Second, 10)
	for i := 0; i < 10; i++ {
		if !l.Acquire("a") {
			t.Fatal("expected true")
		}
	}

	if l.Check("a") {
		t.Fatal("expected false")
	}

	time.Sleep(1 * time.Second)
	if !l.Check("a") {
		t.Fatal("expected true")
	}
	if !l.Acquire("a") {
		t.Fatal("expected true")
	}
}