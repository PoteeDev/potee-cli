package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/PoteeDev/entities/models"
)

func (c *Config) ShowTeams() map[string][]models.EntityInfo {
	url := fmt.Sprintf("%s/api/v1/admin/entities/list", c.Host)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	bearer := fmt.Sprintf("Bearer %s", c.AccessToken)
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
		// for i, entity := range entities["teams"] {
		// 	fmt.Println(i+1, entity.Login, entity.Name, entity.IP, entity.Visible)
		// }
		return entities
	} else {
		log.Println(resp)
	}
	return nil
}

func (c *Config) GenerateVpns() {
	url := fmt.Sprintf("%s/api/v1/admin/entities/generate/vpn", c.Host)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	bearer := fmt.Sprintf("Bearer %s", c.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

}
