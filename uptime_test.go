package uptime

import (
	"testing"
)

func TestBasic(t *testing.T) {
	for i := 0; i < 10; i++ {
		SleepToNextSecond()
		t.Logf("now: %v", ProcUpTime())
	}
}
