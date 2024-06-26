package framework

import (
	"fmt"
	"sync"

	"github.com/DexMus/round_robin_service/biz/constance"
)

var rr *RoundRobin

type RoundRobin struct {
	IPPorts []int
	Length  int
	Cur     int
	lock    *sync.Mutex
}

func InitRoundRobin(ipPorts []int) {
	rr = NewRoundRobin(ipPorts)
}

func (r *RoundRobin) NextRoundRobin(cur int, max int) int {
	return (cur + 1) % max
}

func GetNextRoundRobin() int {
	rr.lock.Lock()
	defer rr.lock.Unlock()
	rr.Cur = rr.NextRoundRobin(rr.Cur, rr.Length)
	return rr.IPPorts[rr.Cur]
}

func GetNextURL() string {
	return fmt.Sprintf(constance.GameInfoURL, GetNextRoundRobin())
}
func NewRoundRobin(ipPorts []int) *RoundRobin {
	return &RoundRobin{
		IPPorts: ipPorts,
		Length:  len(ipPorts),
		Cur:     0,
		lock:    &sync.Mutex{},
	}
}
