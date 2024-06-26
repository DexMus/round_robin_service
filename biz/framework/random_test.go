package framework

import (
	"testing"
	"time"

	"math/rand"

	"github.com/DexMus/round_robin_service/biz/constance"
)

func TestRandLagency(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		result := RandLagency()
		if result < constance.LagencyMin || result > constance.LagencyMax {
			t.Errorf("RandLagency() result out of range: %d", result)
		}
	}
}
