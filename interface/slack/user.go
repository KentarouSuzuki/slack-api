package slack

import (
	"encoding/json"
	"github.com/KentarouSuzuki/slack-api/infrastructure/web_api"
	"log"
)

type UserResponse struct {
	Ok      bool
	Members []User
}

type User struct {
	Id       string
	TeamId   string
	Name     string
	RealName string
	Deleted  bool
	IsAdmin  bool
	IsOwner  bool
	IsBot    bool
	Profile  Profile
}

type Profile struct {
	StatusText  string
	StatusEmoji string
	RealName    string
	DisplayName string
	Email       string
	Team        string
}

func FetchUsers() []User {
	slack := webApi.NewSlackObject()
	url := slack.MakeURL("list", "users")

	var users UserResponse
	jsonb := slack.Send("GET", url, "")
	if err := json.Unmarshal(jsonb, &users); err != nil {
		log.Fatal(err)
	}

	return users.Members
}
