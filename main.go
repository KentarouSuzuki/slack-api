package main

import (
	"fmt"
	"github.com/KentarouSuzuki/slack-member/config"
)

func main() {
	Conf, Err := config.Get()
	fmt.Println(Err)
	fmt.Println(Conf.DataBase.Host)
	fmt.Println(Conf.Slack.Token)
	fmt.Println(Conf.Env.GoEnv)
}