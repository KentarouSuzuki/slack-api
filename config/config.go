package config

import (
	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env			Env
	DataBase	DataBaseConf
	Slack		SlackConf
}

type Env struct {
	GoEnv	string	`default:"local"`
}

type DataBaseConf struct {
	Host	string
	Port	int
	DbName	string
	User	string
	pass	string
}

type SlackConf struct {
	Token	string
}

func Get() (Config, error) {
	var config	Config
	var goenv 	Env
	envconfig.Process("", &goenv)

	if _, err := toml.DecodeFile("./config/" + goenv.GoEnv + ".tml", &config); err != nil {
		return config, err
	}
	config.Env = goenv
	return config, nil
}