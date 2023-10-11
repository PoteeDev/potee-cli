package main

import (
	"github.com/spf13/cobra"
)

func init() {
	adminCmd.AddCommand(usersCmd)
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "register users",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			filename := args[0]
			config.RegisterUsers(filename)
		}

	},
}
