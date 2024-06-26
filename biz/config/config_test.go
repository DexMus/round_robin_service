package config

import (
	"testing"

	"code.byted.org/gopkg/env"
	"github.com/stretchr/testify/assert"
)

type MockEnv struct {
	IsProductFunc func() bool
	IsTestingFunc func() bool
	GetEnvFunc    func(key string) string
}

func (m *MockEnv) IsProduct() bool {
	return m.IsProductFunc()
}

func (m *MockEnv) IsTesting() bool {
	return m.IsTestingFunc()
}

func (m *MockEnv) GetEnv(key string) string {
	return m.GetEnvFunc(key)
}

func TestGetConfInstance(t *testing.T) {
	mockEnv := &MockEnv{
		IsProductFunc: func() bool {
			return false
		},
		IsTestingFunc: func() bool {
			return false
		},
		GetEnvFunc: func(key string) string {
			if key == "RUNTIME_ROOT" {
				return "/tmp"
			}
			return ""
		},
	}
	env.Env = mockEnv

	readFile := func(filename string) ([]byte, error) {
		return []byte(`
			ipports:
			service_ip: "127.0.0.1"
			service_main_port: 8080
			instance_count: 5
			`), nil
	}

	conf := GetConfInstance()
	assert.Equal(t, "127.0.0.1", conf.IPPorts.ServiceIP)
	assert.Equal(t, 8080, conf.IPPorts.ServiceMainPort)
	assert.Equal(t, 5, conf.IPPorts.InstanceCount)
}
