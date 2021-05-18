package models

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	AppName   string `yaml:"appname"`
	HttpPort  uint32 `yaml:"httpport"`
	RunMode   string `yaml:"runmode"`
	SessionOn string `yaml:"sessionon"`
	Mysql     Mysql  `yaml:"mysql"`
	Redis     Redis  `yaml:"redis"`
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
	Db        int  `yaml:"db"`
}

func (c *Config) GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("config/app.yaml")
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Println(err.Error())
	}
	return c
}
