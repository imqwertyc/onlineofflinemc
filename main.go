package main

import "fmt"

func main() {
	// golangci-lint
	fmt.Println("Minecraft Online -> Offline Server Converter\nBy imqwertyc\nGitHub: https://github.com/imqwertyc/offlineonlinemc")
	fmt.Println("")
	fmt.Println("WARNING!! Before you use this tool please make a copy of your server directory to prevent data loss.")
	fmt.Println("WARNING!! Please do NOT make any changes in server files while the tool is working")
	fmt.Println("WARNING!! Please shut down the server before using this tool")
	fmt.Println("This tool is provided AS-IS owner of the tool has no responsibility how it is used or any data loss that may occur. Use at your own risk.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	ConvertOnlineToOffline()
}