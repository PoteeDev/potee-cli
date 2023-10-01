package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (c *Config) ManagerRequest(request string) {
	url := fmt.Sprintf("%s/api/v1/manager/%s", c.Host, request)
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

	var entities map[string]string
	fmt.Println(resp)
	if resp.StatusCode == http.StatusOK {
		json.NewDecoder(resp.Body).Decode(&entities)

	}
}
