package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if steamcmd is installed
	_, err := os.Stat("steamcmd/steamcmd.sh")
	if err != nil {
		//TODO: Check if the codeclysm/extract libary overwrites files
		err = getSteamcmd()
		//fmt.Printf("error: %v\n", err)
		if err != nil {
			fmt.Println("Error Getting Steamcmd")
		}
	}
	//TODO: Implement cli with cobra
	//downloadFromSteamcmd("AppID here", "Workshop Element ID")

}
