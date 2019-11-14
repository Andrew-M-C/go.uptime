package uptime

import (
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	t.Logf("now: %v", ProcUpTime())
	for i := 0; i < 10; i++ {
		SleepToNextSecond()
		t.Logf("now: %v - %v", time.Now(), ProcUpTime())
	}
}
