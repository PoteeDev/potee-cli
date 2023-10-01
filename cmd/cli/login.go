package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to platform",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config)
		if config.Password == "" {
			fmt.Printf("Enter password for %s: ", config.Username)
			fmt.Scanln(&config.Password)
			viper.Set("password", config.Password)
			viper.SafeWriteConfig()
		}
		config.GetToken()
	},
}
