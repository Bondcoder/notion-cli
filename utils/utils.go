package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gookit/color"
)

// Create a Bearer string by appending string access token
var token = "Bearer " + os.Getenv("NOTION_TOKEN")

func MakeRequest(endpoint string) []byte {
	if os.Getenv("NOTION_TOKEN") == "" {
		color.Error.Println("NOTION TOKEN IS NOT SET. READ README.MD TO LEARN HOW TO SET IT UP.")
		os.Exit(1)
	}

	url := "https://api.notion.com/v1/" + endpoint

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", token)
	req.Header.Add("Notion-Version", "2021-05-13")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body
}
