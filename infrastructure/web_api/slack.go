package webApi

import (
	"strings"
	"io/ioutil"
	"net/http"
	"github.com/KentarouSuzuki/slack-api/config"
)

type Slack struct {
	BaseUrl	string
	Token	string
}

func (slack Slack) MakeURL(method string, objects ...string) string {
	url := slack.BaseUrl + "/" + strings.Join(append(objects, method), ".") + "?token=" + slack.Token
	return url
}

func (slack Slack) Send(method string, url string, body string) string {
	client := &http.Client{}

	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	res, _ := client.Do(req)

	byteArr, _ := ioutil.ReadAll(res.Body)
	return string(byteArr)
}

func NewSlackObject() Slack{
	conf, _ := config.Get()
	return Slack{BaseUrl: conf.Slack.BaseUrl, Token: conf.Slack.Token}
}
