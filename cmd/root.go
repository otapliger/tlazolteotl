package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tlazolteotl",
	Short: "Tlazolteotl is a free and open-source system dirty eater",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement TUI
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
