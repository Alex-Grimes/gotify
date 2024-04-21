package gotify

import (
	"sync"
	"time"
)

func truncate(text string, max_length int) string {
	if len(text) > max_length {
		return text[:max_length-3] + "..."
	} else {
		return text
	}
}

func msToHMS(ms int) (int, int, int) {
	seconds := ms / 1000
	hours := seconds / 3600
	seconds -= hours * 3600
	minutes := seconds / 60
	seconds -= minutes * 60
	return hours, minutes, seconds
}

func debounce(s time.Duration) func(func() interface{}) func() interface{} {
	var (
		mu sync.Mutex
		t  time.Time
	)
	return func(f func() interface{}) func() interface{} {
		return func() interface{} {
			mu.Lock()
			defer mu.Unlock()
			t_ := time.Now()
			if t.IsZero() || t_.Sub(t) > s {
				result := f()
				t = time.Now()
				return result
			}
			return nil
		}
	}
}
