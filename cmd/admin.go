package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PoteeDev/entities/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(adminCmd)
	adminCmd.AddCommand(adminTeamsCmd)
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
		ShowTeams()
	},
}

func ShowTeams() {
	url := fmt.Sprintf("%s/api/v1/admin/entities/list", config.Host)
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

	var entities map[string][]models.EntityInfo

	if resp.StatusCode == http.StatusOK {
		json.NewDecoder(resp.Body).Decode(&entities)
		for i, entity := range entities["teams"] {
			fmt.Println(i+1, entity.Login, entity.Name, entity.IP, entity.Visible)
		}
	}
}
