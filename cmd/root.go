package cmd

import (
	"csclub-activities/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "csa",
	Long: "A collection of interactive lessons curated by the Cypress CS Club.\n" +
		"Programmed by Noah Mercedes, president of the 2023-24 year.\n\n" +
		"Note: to get back to this screen at any time, run: `csa help`",
}

func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		// The color package also disables color output if the NO_COLOR environment variable is set to a non-empty string.
		// https://github.com/fatih/color?tab=readme-ov-file#disableenable-colorx

		util.LogErrorAndExit(err)
	}
}
