package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"

	"code.byted.org/gopkg/env"
	"code.byted.org/gopkg/logs"
	"gopkg.in/yaml.v2"
)

var (
	once         sync.Once
	confPath     string
	confInstance *ConfigParam
)

const (
	EnumRuntimeEnvProd = "prod"
	EnumRuntimeEnvBeta = "beta"
	EnumRuntimeEnvDev  = "dev"
)

type ConfigParam struct {
	IPPorts struct {
		ServiceIP       string `yaml:"service_ip"`
		ServiceMainPort int    `yaml:"service_main_port"`
		InstanceCount   int    `yaml:"instance_count"`
	} `yaml:"ip_ports"`
}

func GetConfInstance() *ConfigParam {
	once.Do(func() {
		cp := &ConfigParam{}
		confPath = resolveConfPath()
		logs.Info("[ConfigLoader] config load yaml.(config params=%+v)", cp)
		confInstance = loadYaml(cp)
	})
	return confInstance
}

func resolveConfPath() string {
	// 根据env判断当前环境
	// dev ：开发环境
	// beta：测试环境
	// prod: 生产环境
	runtimeEnv := EnumRuntimeEnvDev
	if env.IsProduct() {
		runtimeEnv = EnumRuntimeEnvProd
	} else if env.IsTesting() {
		runtimeEnv = EnumRuntimeEnvBeta
	}
	runtimeEnv = strings.ToLower(runtimeEnv)

	// bootstrap中设置了RUNTIME_ROOT环境变量
	runtimeRoot := os.Getenv("RUNTIME_ROOT")
	ConfPath := fmt.Sprintf("config.%s.yaml", runtimeEnv)
	return path.Join(runtimeRoot, "conf", ConfPath)
}

func loadYaml(cp *ConfigParam) *ConfigParam {
	var (
		err        error
		byteConfig []byte
	)
	byteConfig, err = ioutil.ReadFile(confPath)
	if err != nil {
		logs.Errorf("[ConfigLoader] read config file failed(filePath=%s, error=%s)", confPath, err)
		panic(err)
	}

	var rawYamlConfig ConfigParam
	err = yaml.Unmarshal(byteConfig, &rawYamlConfig)
	if err != nil {
		logs.Errorf("[ConfigLoader] parse yaml file failed(filePath=%s, error=%s)", confPath, err)
		panic(err)
	}
	return &rawYamlConfig
}
