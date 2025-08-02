package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "LinuxWorkshopDownloader",
	Short: "Linux Workshop Downloader is used to download mods from steam workshop.",
	Long: `Linux Workshop Downloader is used to download mods from steam workshop.
		   source code is available here:
		   github.com/Hydriam/LinuxWorkshopDownloader`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
