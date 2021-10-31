package blocks

import (
	"encoding/json"
	"fmt"
	"notion-cli/utils"
	"strings"
	"time"

	"github.com/gookit/color"
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

const (
	maxLength = 17
)

func ListPages(id string) {
	url := "blocks/" + id + "/children"

	body := utils.MakeRequest(url)
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error while reading the response bytes:", err)
	}

	fmt.Printf("+%s+\n", strings.Repeat("-", 84))
	fmt.Printf("| %20s | %20s | %36s |\n", "TITLE", "LAST UPDATED", "ID")
	for _, rec := range result.Results {
		title := rec.ChildDatabase.Title
		if title != "" {
			if len(title) > maxLength {
				title = title[:maxLength]
			}
			fmt.Printf("|%s|\n", strings.Repeat("-", 84))
			date := rec.LastEditedTime.Format("2006-01-02 15:04:05")
			fmt.Printf("|")
			color.Green.Printf("%22s", title)
			fmt.Printf("| %20s | %30s |\n", date, rec.ID)
		}
	}
	fmt.Printf("+%s+\n", strings.Repeat("-", 84))
}
