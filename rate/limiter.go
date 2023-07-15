package rate

import "time"

type RateLimiter[T comparable] struct {
	Duration time.Duration
	MaxCount    int
	time	 time.Time
	count	 map[T]int
}

func NewLimiter[T comparable](duration time.Duration, maxCount int) *RateLimiter[T] {
	if duration <= 0 {
		panic("duration must be greater than 0")
	}
	if maxCount <= 0 {
		panic("maxCount must be greater than 0")
	}
	return &RateLimiter[T]{
		Duration: duration,
		MaxCount:    maxCount,
		time:   time.Now(),
		count:  make(map[T]int),
	}
}

func (r *RateLimiter[T]) Check(t T) bool {
	if time.Now().Sub(r.time) > r.Duration {
		r.time = time.Now()
		r.count = make(map[T]int)
	}
	val, ok := r.count[t]
	if !ok {
		r.count[t] = 0
		return true
	}
	return val < r.MaxCount
}

func (r *RateLimiter[T]) Acquire(t T) bool {
	if r.Check(t) {
		r.count[t]++
		return true
	}
	return false
}

func (r *RateLimiter[T]) Reset() {
	r.time = time.Now()
	r.count = make(map[T]int)
}
