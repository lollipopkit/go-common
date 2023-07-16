package rate

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

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

// eg: "1/10s" "7/3m"
func FromString[T comparable](s string) (*RateLimiter[T], error) {
	splited := strings.Split(s, "/")
	if len(splited) != 2 {
		return nil, errors.New("invalid format: "+s)
	}
	maxCount, err := strconv.Atoi(splited[0])
	if err != nil {
		return nil, err
	}
	duration, err := time.ParseDuration(splited[1])
	if err != nil {
		return nil, err
	}
	return NewLimiter[T](duration, maxCount), nil
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
