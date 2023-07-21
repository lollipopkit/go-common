package rate

import (
	"sync"
	"time"
)

type Cacher[T any] struct {
	time *time.Time
	duration time.Duration
	data *T
	fn func() (*T, error)
	lock *sync.RWMutex
}

func NewCacher[T any](duration time.Duration, fn func() (*T, error)) *Cacher[T] {
	return &Cacher[T]{
		duration: duration,
		lock: new(sync.RWMutex),
		fn: fn,
	}
}

func (c *Cacher[T]) Get() (*T, error) {
	if c.time == nil || time.Now().Sub(*c.time) > c.duration {
		var err error
		c.lock.Lock()
		c.data, err = c.fn()
		c.lock.Unlock()
		if err != nil {
			return nil, err
		}
		now := time.Now()
		c.time = &now
	}

	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.data, nil
}