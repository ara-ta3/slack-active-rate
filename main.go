package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_WEB_TOKEN"))
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(false)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.PresenceChangeEvent:
				fmt.Println(ev.Presence)
				fmt.Println(ev.User)
				fmt.Println(time.Now().Format("2006-01-02 15:00:00"))
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			default:
			}
		}
	}
}
