package config

import "github.com/BurntSushi/toml"

type Config struct {
	DataBase	DataBaseConf
	Slack		SlackConf
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
	var config Config

	if _, err := toml.DecodeFile("./config/local.tml", &config); err != nil {
		return config, err
	}
	return config, nil
}