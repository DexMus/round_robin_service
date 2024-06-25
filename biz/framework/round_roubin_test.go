package framework

import (
	"testing"
)

func TestGetNextRoundRobin(t *testing.T) {
	InitRoundRobin([]int{8080, 8081, 8082})
	for i := 0; i < 10; i++ {
		ipPort := GetNextRoundRobin()
		if ipPort != rr.IPPorts[rr.Cur] {
			t.Errorf("Expected %d, got %d", rr.IPPorts[rr.Cur], ipPort)
		}
	}
}
