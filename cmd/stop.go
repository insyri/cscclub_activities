package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"csclub-activities/util"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Clears the lesson checkpoint",
	Aliases: []string{"clear", "new"},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := util.NewSave()
		if err != nil {
			return err
		}
		fmt.Println("Checkpoint cleared")
		return nil
	},
}
