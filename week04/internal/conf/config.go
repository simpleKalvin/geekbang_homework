package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"reflect"
	"runtime"
	"sync/atomic"
)

var (
	_ Config = (*config)(nil)
)

type Config interface {
	Load() error
	Values(key string) string
}

type values struct {
	atomic.Value
}

type config struct {
	values
	Server server `yaml:"server"`
}


type server struct {
	HttpServer httpServer `yaml:"http"`
}

type httpServer struct {
	Host string `yaml:"addr"`
	Port string	`yaml:"port"`
	ReadTimeout string `yaml:"read_timeout"`
	WriteTimeout string `yaml:"write_timeout"`
}

func (this *config) Load() error {
	// load the file
	yamlPath := "configs/config.yaml"
	yamlPath = getCurrentAbPathByCaller() +"/../../"+ yamlPath
	yamlConfig, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		panic(yamlPath+" not found")
	}
	err = yaml.Unmarshal(yamlConfig, this)
	//values := reflect.ValueOf(this.Server.HttpServer)
	this.values.Store(this.Server.HttpServer)
	if err != nil {
		panic("unmarshal failed")
	}
	return nil
}

func (this *config) Values(key string) string {
	// 通过atomic获取数据
	dataValues := this.values.Load().(httpServer)
	// reflection
	refValue := reflect.ValueOf(dataValues)
	return refValue.FieldByName(key).String()
}

func NewConfig() Config {
	appConfig := config{}
	appConfig.Load()
	return &appConfig
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}