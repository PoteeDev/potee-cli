package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GeneratePassword(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Subnet   string `json:"subnet,omitempty"`
	Ip       string `json:"ip,omitempty"`
	Visible  bool   `json:"visible,omitempty"`
}

func (c *Config) RegisterUsers(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	users := []User{}

	if err = json.Unmarshal([]byte(file), &users); err != nil {
		log.Fatalln(err)
	}

	for i, user := range users {
		user.Ip = fmt.Sprintf("10.0.%d.10", i+1)
		user.Subnet = fmt.Sprintf("10.0.%d.0/24", i+1)
		user.Password = GeneratePassword(24)
		user.Visible = true
		users[i] = user
	}
	fmt.Println(users)

	file, err = json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	fullFilename := fmt.Sprintf("full-%s", filename)
	if _, err := os.Stat(fullFilename); errors.Is(err, os.ErrNotExist) {
		if err = os.WriteFile(fullFilename, file, 0644); err != nil {
			log.Fatalln(err)
		}
	}

	url := fmt.Sprintf("%s/api/v1/admin/entities/registration", c.Host)

	for _, user := range users {
		jsonPayload := new(bytes.Buffer)
		json.NewEncoder(jsonPayload).Encode(user)

		req, err := http.NewRequest(http.MethodPost, url, jsonPayload)
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
}
