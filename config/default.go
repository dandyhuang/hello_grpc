package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type HelloConfig struct {
	Server struct {
		Addr string `yaml:"address"`
		Port string `yaml:"port"`
	}
	Mysql struct {
		User     string `yaml:"user"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	}
}

func newHelloConfig() *HelloConfig {
	return &HelloConfig{}
}

// DefaultConfigLoader 默认配置加载器
var DefaultConfig = newHelloConfig()

func (cfg HelloConfig) load(path string) (*HelloConfig, error) {
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(DefaultConfig)
	}
	return DefaultConfig, nil
}
