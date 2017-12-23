package main

import (
	"io/ioutil"

	"github.com/nlopes/slack"
	yaml "gopkg.in/yaml.v1"
)

var (
	botID   string
	botName string
)

type Bot struct {
	api *slack.Client
	rtm *slack.RTM
}

type Token struct {
	Id string
}

func NewBot(token string) *Bot {
	bot := new(Bot)
	bot.api = slack.New(token)
	bot.rtm = bot.api.NewRTM()
	return bot
}

func GetToken() (string, error) {
	buf, err := ioutil.ReadFile("token.yml")
	if err != nil {
		return "", err
	}

	var token Token
	err = yaml.Unmarshal(buf, &token)
	if err != nil {
		return "", err
	}

	return token.Id, nil
}

func main() {
	token, err := GetToken()
	if err != nil {
		panic(err)
	}

	bot := NewBot(token)

	go bot.rtm.ManageConnection()

	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				botID = ev.Info.User.ID
				botName = ev.Info.User.Name

			case *slack.TeamJoinEvent:
				params := slack.NewPostMessageParameters()
				params.AsUser = true
				bot.api.PostMessage("@"+ev.User.ID, "hoge", params)
			}
		}
	}
}
