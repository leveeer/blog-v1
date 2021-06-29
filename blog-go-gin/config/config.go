package config

import (
	"blog-go-gin/logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var OnceDo = sync.Once{}
var config = &Config{}

type Config struct {
	AppName   string `yaml:"appname"`
	HttpPort  uint32 `yaml:"httpport"`
	RunMode   string `yaml:"runmode"`
	SessionOn string `yaml:"sessionon"`
	//LogLevel  string `yaml:"loglevel"`
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	Ws    WS    `yaml:"ws"`
}

type Mysql struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint32 `yaml:"port"`
	DbName   string `yaml:"dbName"`
}

type Redis struct {
	RedisConn string `yaml:"redisconn"`
	RedisPwd  string `yaml:"redispwd"`
	Db        int    `yaml:"db"`
}

type WS struct {
	ID    uint32 `yaml:"id"`
	Host  string `yaml:"host"`
	Port  uint32 `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

func (c *Config) loadConf() {
	yamlFile, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
	config = c
}

func GetConf() *Config {
	return config
}

func init() {
	OnceDo.Do(func() {
		config.loadConf()
	})
}
