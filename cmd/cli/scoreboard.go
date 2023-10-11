package main

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scoreCmd)
}

var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "Use admin functions",
	Run: func(cmd *cobra.Command, args []string) {
		config.GetScoreboard()
	},
}
