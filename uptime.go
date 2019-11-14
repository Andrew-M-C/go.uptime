package uptime

import (
	"strconv"
	"time"
)

// ProcUpTime returns durations that this process had ran
func ProcUpTime() time.Duration {
	now := time.Now()
	seconds, nanoseconds := findUpTime(now.String())
	return time.Duration(seconds)*time.Second + time.Duration(nanoseconds)*time.Nanosecond
}

func findUpTime(t string) (int, int) {
	l := len(t)
	nanoseconds := t[l-9:]

	i := l - 11
	for ; t[i] != '+'; i-- {
		// nothing
	}
	seconds := t[i+1 : l-10]

	sec, _ := strconv.Atoi(seconds)
	nsec, _ := strconv.Atoi(nanoseconds)
	return sec, nsec
}

func findUpTimeNanoPart(t string) int {
	l := len(t)
	nanoseconds := t[l-9:]
	nsec, _ := strconv.Atoi(nanoseconds)
	return nsec
}

// SleepToNextSecond sleeps to next proc up second
func SleepToNextSecond() {
	now := time.Now()
	nanoseconds := findUpTimeNanoPart(now.String())
	time.Sleep(time.Second - time.Duration(nanoseconds))
}
