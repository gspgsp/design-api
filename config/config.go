package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var (
	dir = flag.String("dir", "D:/gopath/src/design-api", "请输入项目根目录: -dir xxxxx")

	//解析yaml配置文件
	Config = parseYaml()
)

func parseYaml() *configuration {
	flag.Parse()
	config := new(configuration)

	config, err := config.yaml(*dir + "/config/config.yaml")

	if err != nil {
		log.Panic(err)
	}

	log.Printf("服务配置完成")
	return config
}

type configuration struct {
	Mysql mysql `json:"mysql",yaml:"mysql"`
	Redis redis `json:"redis",yaml:"redis"`
	Addr  addr  `json:"addr",yaml:"addr"`
}

func (conf *configuration) yaml(dir string) (*configuration, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil { //file 实现了 Read方法
		return nil, err
	}

	err = yaml.UnmarshalStrict(data, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

type mysql struct {
	DbConnect  string `yaml:"db_connect",json:"db_connect"`
	DbHost     string `yaml:"db_host",json:"db_host"`
	DbPort     string `yaml:"db_port",json:"db_port"`
	DbDatabase string `yaml:"db_database",json:"db_database"`
	DbUsername string `yaml:"db_username",json:"db_username"`
	DbPassword string `yaml:"db_password",json:"db_password"`
	DbPrefix   string `yaml:"db_prefix",json:"db_prefix"`
}

type redis struct {
	RedisHost     string `yaml:"redis_host",json:"redis_host"`
	RedisPort     string `yaml:"redis_port",json:"redis_port"`
	RedisDatabase int    `yaml:"redis_database",json:"redis_database"`
	RedisUsername string `yaml:"redis_username",json:"redis_username"`
	RedisPassword string `yaml:"redis_password",json:"redis_password"`
}

type addr struct {
	Tcp  string `yaml:"tcp",json:"tcp"`
	Unix string `yaml:"unix",json:"unix"`
}
