package main

import (
	"fmt"
	"github.com/KentarouSuzuki/slack-api/infrastructure/web_api"
)

func main() {
	slack := webApi.NewSlackObject()

	url := slack.MakeURL("list", "users")
	fmt.Println(slack.Send("GET", url, ""))
}