package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
		ManagerRequest("start")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Show teams",
	Run: func(cmd *cobra.Command, args []string) {
		ManagerRequest("stop")
	},
}

func ManagerRequest(request string) {
	url := fmt.Sprintf("%s/api/v1/manager/%s", config.Host, request)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	bearer := fmt.Sprintf("Bearer %s", config.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var entities map[string]string
	fmt.Println(resp)
	if resp.StatusCode == http.StatusOK {
		json.NewDecoder(resp.Body).Decode(&entities)

	}
}
