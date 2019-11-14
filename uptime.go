package uptime

import (
	"time"
)

var (
	_uptime time.Time
)

func init() {
	_uptime = time.Now()
}

// ProcUpTime returns durations that this process had ran
func ProcUpTime() time.Duration {
	return time.Now().Sub(_uptime)
}

// SleepToNextSecond sleeps to next proc up second
func SleepToNextSecond() {
	now := time.Now()
	time.Sleep(time.Second - time.Duration(now.Nanosecond())*time.Nanosecond)
}
