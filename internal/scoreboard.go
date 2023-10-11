package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func (c *Config) GetScoreboard() {
	url := fmt.Sprintf("%s/api/v1/scores/user?name=naliway", c.Host)
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
	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println(err.Error())
	}
	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON", err)
	}
	fmt.Println(result)
}
