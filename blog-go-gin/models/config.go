package models

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	AppName   string `yaml:"appname"`
	HttpPort  uint32 `yaml:"httpport"`
	RunMode   string `yaml:"runmode"`
	SessionOn string `yaml:"sessionon"`
	Mysql     Mysql  `yaml:"mysql"`
}

type Mysql struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint32 `yaml:"port"`
	DbName   string `yaml:"dbName"`
}

func (c *Config) GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("config/app.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
