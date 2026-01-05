package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/sqweek/dialog"
)

func GetOfflineUuid(player string) string {
	return uuid.NewMD5(uuid.Nil, []byte("OfflinePlayer:"+player)).String()
}

func LookupOnlineUsernane(uuid string) (string, error) {
	res, err := http.Get("https://api.minecraftservices.com/minecraft/profile/lookup/" + uuid)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to lookup username for UUID %s: status code %d", uuid, res.StatusCode)
	}
	type Response struct {
		Id	   string `json:"id"`
		Name   string `json:"name"`
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	return response.Name, nil
}

func GetServerPath() string {
	path, err := dialog.Directory().Title("Select Minecraft Server Directory").Browse()
	if err != nil {
		panic(err)
	}
	return path
}

type Option struct {
	Label string
	Value int
}

func ChooseOption(options []Option) int {
	// Prepare labels array
	labels := make([]string, len(options))
	for i, option := range options {
		labels[i] = option.Label
	}
	// Prepare values array
	values := make([]int, len(options))
	for i, option := range options {
		values[i] = option.Value
	}
	var opt int
	// Prompt user until valid input is received
	for {
		// Show options
		fmt.Println("Please choose an option and press Enter:")
		for i, label := range labels {
			fmt.Printf("[%d] %s\n", i+1, label)
		}
		// Read user input
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n') // Wait for user to press Enter
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		// Parse user input
		var choice int
		_, err = fmt.Sscanf(input, "%d", &choice)
		if err != nil || choice < 1 || choice > len(options) {
			fmt.Println("Invalid choice, please try again.")
			continue
		}
		opt = values[choice-1]
		return opt
	}
}