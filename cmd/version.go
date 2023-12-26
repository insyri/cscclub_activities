package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// TODO: add specific info (release, commit hash, platform, if update available)

const VERSION string = "v0.0.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the program",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}
