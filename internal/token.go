package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func (c *Config) GetToken() {
	values := map[string]string{"username": c.Username, "password": c.Password}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("%s/api/v1/auth/login", c.Host)
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
