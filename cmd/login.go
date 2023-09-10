package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
		getToken()
	},
}

func getToken() {
	values := map[string]string{"username": config.Username, "password": config.Password}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("%s/api/v1/auth/login", config.Host)
	fmt.Println(url)
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	if resp.StatusCode == http.StatusOK {
		viper.Set("access_token", res["access_token"])
		viper.WriteConfig()
	}
}
