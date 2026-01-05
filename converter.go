package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ConvertOnlineToOffline() {
	fmt.Println("Converting from Online to Offline...")
	serverDir := GetServerPath()
	advancementsPath := path.Join(serverDir, "world", "advancements")
	playerdataPath := path.Join(serverDir, "world", "playerdata")
	serverPropertiesPath := path.Join(serverDir, "server.properties")
	// Stat the directories to ensure they exist
	aStat, err := os.Stat(advancementsPath)
	pdStat, err := os.Stat(playerdataPath)
	spStat, err := os.Stat(serverPropertiesPath)
	if err != nil {
		panic(err)
	}
	if !aStat.IsDir() || !pdStat.IsDir() || spStat.IsDir() {
		panic("Advancements or Playerdata path is not a directory or server.properties is a directory")
	}
	// Change online-mode to false in server.properties
	fmt.Println("Changing online-mode property to false in server.properties...")
	serverProperties, err := os.ReadFile(serverPropertiesPath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(serverProperties), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "online-mode=") {
			lines[i] = "online-mode=false"
		}
	}
	newServerProperties := strings.Join(lines, "\n")
	err = os.WriteFile(serverPropertiesPath, []byte(newServerProperties), spStat.Mode())
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	

	// Convert playerdata files from online UUIDs to offline UUIDs
	fmt.Println("Converting playerdata files from online UUIDs to offline UUIDs...")
	playerFiles, err := filepath.Glob(path.Join(playerdataPath, "*.dat"))
	if err != nil {
		panic(err)
	}
	for _, playerFile := range playerFiles {
		base := filepath.Base(playerFile)
		onlineUUID := strings.TrimSuffix(base, ".dat")
		onlineUsername, err := LookupOnlineUsernane(onlineUUID)
		if err != nil {
			fmt.Printf("Failed to lookup username for UUID %s: %v\n", onlineUUID, err)
			continue
		}
		offlineUUID := GetOfflineUuid(onlineUsername)
		newPlayerFile := path.Join(playerdataPath, offlineUUID+".dat")
		err = os.Rename(playerFile, newPlayerFile)
		if (err != nil) {
			fmt.Printf("Failed to rename playerdata file for UUID %s: %v\n", onlineUUID, err)
			continue
		}
		fmt.Println("  Migrated user:", onlineUsername)
	}
	fmt.Println("Done!")
	// Convert advancements files from online UUIDs to offline UUIDs
	fmt.Println("Converting advancements files from online UUIDs to offline UUIDs...")
	advancementFiles, err := filepath.Glob(path.Join(advancementsPath, "*.json"))
	if err != nil {
		panic(err)
	}
	for _, advancementFile := range advancementFiles {
		base := filepath.Base(advancementFile)
		onlineUUID := strings.TrimSuffix(base, ".json")
		onlineUsername, err := LookupOnlineUsernane(onlineUUID)
		if err != nil {
			fmt.Printf("Failed to lookup username for UUID %s: %v\n", onlineUUID, err)
			continue
		}
		offlineUUID := GetOfflineUuid(onlineUsername)
		newAdvancementFile := path.Join(advancementsPath, offlineUUID+".json")
		err = os.Rename(advancementFile, newAdvancementFile)
		if (err != nil) {
			fmt.Printf("Failed to rename advancement file for UUID %s: %v\n", onlineUUID, err)
			continue
		}
		fmt.Println("  Migrated user:", onlineUsername)
	}
	fmt.Println("Done!")
	fmt.Println("Conversion complete!")
	fmt.Println("Thanks for using my tool! Have fun playing!")
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
	os.Exit(0)
}
