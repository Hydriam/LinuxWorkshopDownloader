package libLWD

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/codeclysm/extract/v4"
)

func GetSteamcmd() error {
	//Thanks for https://gist.github.com/cnu/026744b1e86c6d9e22313d06cba4c2e9

	//Download the archive with steamcmd binaries
	fmt.Println("Downloading steamcmd.")

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

	fmt.Println("Steamcmd Downloaded.")
	//Extract the archive
	fmt.Println("Extracting Steamcmd.")

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
	fmt.Println("Downloaded Steamcmd Succesfully.")
	return nil
}

func DownloadFromSteamcmd(appID string, workshopIDs []string, debug bool) error {
	fmt.Println("Downloading from steamcmd.")
	// cmdt = command template
	cmdt := []string{
		"./steamcmd/steamcmd.sh",
		"+login", "anonymous",
	}
	for i := 0; i < len(workshopIDs); i++ {
		cmdt = append(cmdt, "+workshop_download_item", appID, workshopIDs[i])
	}
	cmdt = append(cmdt, "+quit")
	cmd := exec.Command(cmdt[0], cmdt[1:]...) //cmdt[1:]... makes it use every object in array
	// Debug Mode
	if debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Steamcmd Failed.")
		fmt.Println("Make sure that you have 32bit version of glibc installed on your system.")
		return err
	}
	fmt.Println("The file has been downloaded.")
	fmt.Println("It should be under ~/.local/share/Steam/steamapps/workshop/")
	return nil
}
