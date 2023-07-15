package rate

import "time"

type RateLimiter[T comparable] struct {
	Duration time.Duration
	MaxCount    int
	time	 time.Time
	count	 map[T]int
}

func NewRateLimiter[T comparable](duration time.Duration, maxCount int) *RateLimiter[T] {
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

func (r *RateLimiter[T]) Can(t T) bool {
	if time.Now().Sub(r.time) > r.Duration {
		r.time = time.Now()
		r.count = make(map[T]int)
	}
	if r.count[t] < r.MaxCount {
		if val, ok := r.count[t]; ok {
			r.count[t] = val + 1
		} else {
			r.count[t] = 1
		}
		return true
	}
	return false
}

func (r *RateLimiter[T]) Reset() {
	r.time = time.Now()
	r.count = make(map[T]int)
}
