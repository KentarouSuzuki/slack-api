package slack

import (
	"encoding/json"
	"github.com/KentarouSuzuki/slack-api/infrastructure/web_api"
	"log"
)

type UserGroupResponse struct {
	Ok         bool
	Usergroups []UserGroup
}

type UserGroup struct {
	Id     string
	TeamId string
	Name   string
}

func FetchUserGroups() []UserGroup {
	slack := webApi.NewSlackObject()
	url := slack.MakeURL("list", "usergroups")

	var res UserGroupResponse
	jsonb := slack.Send("GET", url, "")
	if err := json.Unmarshal(jsonb, &res); err != nil {
		log.Fatal(err)
	}

	return res.Usergroups
}
