package conf

import (
	"io/ioutil"
	"github.com/go-yaml/yaml"
)

type Config struct {
	Servers     []map[string]string
	Sleeptime   int
	Maxidle     int
	Maxactive   int
	Idletimeout int
	Datatype    string
	Datapath    string
	Logpath     string
	Logname     string
	Loglevel    int
	Serverip    string
	Serverport  int
	Staticdir   string
	Tpldir      string
	Debugmode   int
}

type MyConfig struct {
	Name string
	Age  int
	MySex int
}

func NewConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := new(Config)
	if err := yaml.Unmarshal([]byte(data), config); err != nil {
		return nil, err
	}
	return config, nil
}

func NewConfig1(path string) (*MyConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	myconfig := new(MyConfig)
	if err := yaml.Unmarshal([]byte(data), myconfig); err != nil {
		return nil, err
	}
	return myconfig, nil
}
