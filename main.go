package main

import (
	"flag"
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
	slackToken := flag.String("t", "default", "A personal Slack Token")
	flag.Parse()

	api := slack.New(*slackToken)
	users, err := api.GetUsers()

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	if len(users) == 0 {
		fmt.Println("No users found!")
	}

	fmt.Println("Deleted users:")

	for _, user := range users {
		if user.Deleted {
			fmt.Println("-", fmt.Sprintf("%s(%s)", user.Name, user.RealName), "updated", user.Updated)
		}
	}
}
