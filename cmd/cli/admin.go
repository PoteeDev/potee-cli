package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(adminCmd)
	adminCmd.AddCommand(adminTeamsCmd)
	adminCmd.AddCommand(adminGenerateVpnsCmd)
}

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Use admin functions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var adminTeamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Show teams",
	Run: func(cmd *cobra.Command, args []string) {
		teams := config.ShowTeams()
		for id, team := range teams {
			if team.Visible {
				fmt.Println(id+1, team.Login, team.Name, team.IP)
			}
		}
	},
}

var adminGenerateVpnsCmd = &cobra.Command{
	Use:   "vpn",
	Short: "Generate vpn config for teams",
	Run: func(cmd *cobra.Command, args []string) {
		config.GenerateVpns()
	},
}
