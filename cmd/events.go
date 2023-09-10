package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"golang.org/x/net/websocket"
)

type Message struct {
	Id      int         `json:"id,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

type RoundInfo struct {
	TeamName string              `json:"team_name,omitempty"`
	TeamHost string              `json:"team_host,omitempty"`
	Services map[string]Services `json:"services,omitempty"`
}

type Services struct {
	PingStatus int
	Checkers   map[string]Checker
	Exploits   map[string]Exploit //exploit name and status
}

type Checker struct {
	GetStatus int
	PutStatus int
}

type Exploit struct {
	Cost   int
	Status int
}

func init() {
	rootCmd.AddCommand(eventsCmd)
}

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "show events",
	Run: func(cmd *cobra.Command, args []string) {
		readEvents()
	},
}

// Client.
func readEvents() {

	ws, err := websocket.Dial(
		fmt.Sprintf("ws://%s/ws", config.Host),
		"",
		fmt.Sprintf("%s/ws", config.Host))
	for {
		var m Message
		// err = websocket.JSON.Receive(ws, &m)
		//var m interface{}
		err = websocket.JSON.Receive(ws, &m)
		if err != nil {
			if err == io.EOF {
				break
			}
			//fmt.Printf("Received: %+v\n", m)
			// log.Fatalln(err)
		}
		fmt.Printf("Received: %+v\n", m.Message)
	}
}
