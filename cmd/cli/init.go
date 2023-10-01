package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init cli tool",
	Run: func(cmd *cobra.Command, args []string) {
		viper.SafeWriteConfig()
	},
}
