package common

import "time"

func MeasureUS(fn func()) int64 {
	start := time.Now()
	fn()
	return time.Since(start).Microseconds()
}
