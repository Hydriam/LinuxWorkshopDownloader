package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/codeclysm/extract/v4"
)

func getSteamcmd() error {
	//Thanks for https://gist.github.com/cnu/026744b1e86c6d9e22313d06cba4c2e9

	//Download the archive with steamcmd binaries
	fmt.Println("Downloading steamcmd")

	out, err := os.Create("steamcmd.tar.gz")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get("https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Steamcmd Downloaded")
	//Extract the archive
	fmt.Println("Extracting Steamcmd")

	_, err = os.Stat("steamcmd")
	if os.IsNotExist(err) {
		err = os.Mkdir("steamcmd", 0755)
		if err != nil {
			return err
		}
	}

	file, err := os.Open("steamcmd.tar.gz")
	if err != nil {
		return err
	}
	err = extract.Gz(context.TODO(), file, "steamcmd", nil)
	if err != nil {
		defer file.Close()
		defer os.Remove("steamcmd.tar.gz")
		return err
	}
	defer file.Close()
	os.Remove("steamcmd.tar.gz")
	fmt.Println("Downloaded Steamcmd Succesfully")
	return nil
}

// workshopID = ID of workshop mod
// appID = appID of game the mod belongs to
func downloadFromSteamcmd(appID string, workshopID string) error {
	//TODO : implement downloading multiple element in one instance of steamcmd.sh
	// The command
	cmd := exec.Command(
		"./steamcmd/steamcmd.sh",
		"+login", "anonymous",
		"+workshop_download_item", appID, workshopID,

		"+quit",
	)
	// Redirect command output to our output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Steamcmd Failed.")
		return err
	}
	fmt.Println("The file has been downloaded")
	return nil
}
