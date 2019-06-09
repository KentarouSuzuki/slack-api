package main

import (
	"fmt"
	"github.com/KentarouSuzuki/slack-api/interface/slack"
)

func main() {
	users := slack.FetchUsers()
	for _, user := range users {
		fmt.Println(user.Name)
	}
}
