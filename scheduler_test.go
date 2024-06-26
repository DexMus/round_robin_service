package main

import (
	"context"
	"testing"

	"github.com/DexMus/round_robin_service/biz/framework"
)

type MockFramework struct {
	GetFreePortsFunc   func(count int) ([]int, error)
	InitRoundRobinFunc func(ports []int)
}

func (m *MockFramework) GetFreePorts(count int) ([]int, error) {
	return m.GetFreePortsFunc(count)
}

func (m *MockFramework) InitRoundRobin(ports []int) {
	m.InitRoundRobinFunc(ports)
}

func TestStartAndKeepSimpleAPI(t *testing.T) {
	mockFrame := &MockFramework{
		GetFreePortsFunc: func(count int) ([]int, error) {
			return []int{8081, 8082}, nil
		},
		InitRoundRobinFunc: func(ports []int) {},
	}
	framework.Framework = mockFrame

	ctx := context.Background()
	count := 2
	serviceIP := "127.0.0.1"

	startAndKeepSimpleAPI(ctx, count, serviceIP)

}
