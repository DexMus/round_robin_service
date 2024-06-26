package framework

import (
	"fmt"
	"testing"

	"github.com/DexMus/round_robin_service/biz/constance"
)

func TestGetNextRoundRobin(t *testing.T) {
	// 初始化测试数据
	rr = &RoundRobin{
		Cur:     0,
		Length:  5,
		IPPorts: []int{10, 20, 30, 40, 50},
	}

	// 第一次调用函数，预期返回 10
	expected := 10
	result := GetNextRoundRobin()
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// 第二次调用函数，预期返回 20
	expected = 20
	result = GetNextRoundRobin()
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestGetNextURL(t *testing.T) {
	expectedURL := fmt.Sprintf(constance.GameInfoURL, 1)
	result := GetNextURL()
	if result != expectedURL {
		t.Errorf("Expected URL: %s, but got: %s", expectedURL, result)
	}
}
