package main

import (
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"sort"
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
		return
	}

	fmt.Println("Deleted users:")

	sort.Slice(users, func(i, j int) bool {
		return users[i].Updated < users[j].Updated
	})

	for _, user := range users {
		updatedTime := user.Updated.Time()
		if user.Deleted {
			fmt.Print(fmt.Sprintf("- %s ", user.Name))
			if user.RealName != "" {
				fmt.Print(fmt.Sprint("(%s)", user.RealName))
			}
			fmt.Println("on", updatedTime.Format("2006/01/02"))
		}
	}
}
