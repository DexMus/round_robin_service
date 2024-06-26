package framework

import (
	"math/rand"

	"github.com/DexMus/round_robin_service/biz/constance"
)

func RandLagency() int {
	return constance.LagencyMin + rand.Intn(constance.LagencyMax-constance.LagencyMin)
}
