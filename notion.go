package main

import (
	"flag"
	"os"
	"time"

	"notion-cli/blocks"
	"notion-cli/users"
)

type Response struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		ID             string    `json:"id"`
		CreatedTime    time.Time `json:"created_time"`
		LastEditedTime time.Time `json:"last_edited_time"`
		HasChildren    bool      `json:"has_children"`
		Archived       bool      `json:"archived"`
		Type           string    `json:"type"`
		ChildDatabase  struct {
			Title string `json:"title"`
		} `json:"child_database,omitempty"`
		ChildPage struct {
			Title string `json:"title"`
		} `json:"child_page,omitempty"`
	} `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}

func main() {
	listCommand := flag.NewFlagSet("ls", flag.ExitOnError)
	idOpt := listCommand.String("notionId", "", "Page, database or block ID that you want to query")

	userCommand := flag.NewFlagSet("users", flag.ExitOnError)

	flag.Parse()
	switch os.Args[1] {
	case "ls":
		listCommand.Parse((os.Args[2:]))
		blocks.ListPages(*idOpt)
	case "users":
		userCommand.Parse((os.Args[2:]))
		if len(os.Args) < 3 {
			users.ListUsers()
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
