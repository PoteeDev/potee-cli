package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(manageCmd)
	manageCmd.AddCommand(startCmd)
	manageCmd.AddCommand(stopCmd)
}

var manageCmd = &cobra.Command{
	Use:   "manager",
	Short: "Use admin functions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Show teams",
	Run: func(cmd *cobra.Command, args []string) {
		config.ManagerRequest("start")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Show teams",
	Run: func(cmd *cobra.Command, args []string) {
		config.ManagerRequest("stop")
	},
}
