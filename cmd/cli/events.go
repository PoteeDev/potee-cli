package main

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(eventsCmd)
}

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "show events",
	Run: func(cmd *cobra.Command, args []string) {
		config.ReadEvents()
	},
}
