package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	libLWD "github.com/Hydriam/LinuxWorkshopDownloader/LibLWD"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads object from steam workshop",
	Long: `Downloads object from steam workshop.
Pass AppId of game the mod belongs to as first argument.
All other arguments should be ids of mods.`,
	Run: func(cmd *cobra.Command, args []string) {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			debug = false
		}
		err = libLWD.DownloadFromSteamcmd(args[0], args[1:], debug)
		if err != nil {
			fmt.Println("Run with -d flag to see whats the problem")
		}
	},
}

func init() {
	// Check if steamcmd is installed
	_, err := os.Stat("steamcmd/steamcmd.sh")
	if err != nil {
		println("Couldn't find steamcmd, installing it now.")
		//TODO: Check if the codeclysm/extract libary overwrites files
		err = libLWD.GetSteamcmd()
		//fmt.Printf("error: %v\n", err)
		if err != nil {
			fmt.Println("Error Getting Steamcmd.")
			fmt.Println("Please Get It Manually.")
		}
	}
	downloadCmd.Flags().BoolP("debug", "d", false, "prints steamcmd output")
	rootCmd.AddCommand(downloadCmd)
}
