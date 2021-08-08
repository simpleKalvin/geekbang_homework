package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync/atomic"
)

var (
	_ Config = (*config)(nil)
)

type Config interface {
	Load() error
	Values(key string) values
}

type values struct {
	atomic.Value
}

type config struct {
	httpServer
}

type httpServer struct {
	host string `yaml:"addr"`
	port string	`yaml:"port"`
	readTime string `yaml:"read_time"`
	writeTime string `yaml:"write_time"`
}

func (this *config) Load() error {
	// load the file
	yamlPath := "../configs/config.yaml"
	yamlConfig, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		panic(yamlPath+" not found")
	}
	err = yaml.Unmarshal(yamlConfig, this.httpServer)
	if err != nil {
		panic("unmarshal failed")
	}
	fmt.Println(this.httpServer.port)
	return nil
}

func (this *config) Values(key string) values {
	fmt.Println(this.httpServer)
	return values{}
}

func NewConfig() Config {
	appConfig := config{}
	appConfig.Load()
	return &appConfig
}