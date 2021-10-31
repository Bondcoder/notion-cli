package users

import (
	"encoding/json"
	"fmt"
	"notion-cli/utils"
	"strings"

	"github.com/gookit/color"
)

type Response struct {
	Object  string `json:"object"`
	Results []struct {
		Object    string      `json:"object"`
		ID        string      `json:"id"`
		Name      string      `json:"name"`
		AvatarURL interface{} `json:"avatar_url"`
		Type      string      `json:"type"`
		Person    struct {
			Email string `json:"email"`
		} `json:"person,omitempty"`
		Bot struct {
			Owner struct {
				Type      string `json:"type"`
				Workspace bool   `json:"workspace"`
			} `json:"owner"`
		} `json:"bot,omitempty"`
	} `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}

const (
	maxLength = 17
)

func ListUsers() {
	url := "users"

	body := utils.MakeRequest(url)
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error while reading the response bytes:", err)
	}

	fmt.Printf("+%s+\n", strings.Repeat("-", 84))
	fmt.Printf("| %20s | %20s | %36s |\n", "NAME", "TYPE", "ID")
	for _, rec := range result.Results {
		title := rec.Name
			if len(title) > maxLength {
				title = title[:maxLength]
			}
		
		fmt.Printf("|%s|\n", strings.Repeat("-", 84))
		fmt.Printf("|")
		color.Green.Printf("%22s", title)
		fmt.Printf("| %20s | %30s |\n", rec.Type, rec.ID)

	}
	fmt.Printf("+%s+\n", strings.Repeat("-", 84))
}
