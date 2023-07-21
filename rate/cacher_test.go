package rate_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/lollipopkit/gommon/rate"
)

var n = 0

func TestCacher(t *testing.T) {
	fn := func() (*int, error) {
		time.Sleep(time.Millisecond * 200)
		n = rand.Intn(160)
		return &n, nil
	}
	c := rate.NewCacher[int](time.Second, fn)

	_cacherGet(t, c)
	for i := 0; i < 5; i++ {
		go _cacherGet(t, c)
	}
	time.Sleep(time.Second)
	_cacherGet(t, c)
}

func _cacherGet(t *testing.T, c *rate.Cacher[int]) {
	b, _ := c.Get()
	t.Logf("%d", *b)
	if *b != n {
		t.Error("expected", n, ", got", *b)
	}
}