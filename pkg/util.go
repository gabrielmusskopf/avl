package avl

import "time"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func measure(f func(*int)) (int, time.Duration) {
	iter := 0
	start := time.Now()
	f(&iter)
	return iter, time.Since(start)
}
