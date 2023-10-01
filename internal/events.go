package internal

import (
	"fmt"
	"io"
	"log"

	"github.com/PoteeDev/events-stream/models"
	"golang.org/x/net/websocket"
)

type Message struct {
	Id      int                `json:"id,omitempty"`
	Message models.TeamsEvents `json:"message,omitempty"`
}

// Client.
func (c *Config) ReadEvents() {

	ws, err := websocket.Dial(
		fmt.Sprintf("ws://%s/ws", c.Host),
		"",
		fmt.Sprintf("http://%s", c.Host),
	)
	if err != nil {
		log.Fatalln(err)
	}
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
		for _, team := range m.Message.Teams {
			var servicesInfo []string
			for _, service := range team.Services {
				servicesInfo = append(servicesInfo, fmt.Sprintf("%s: %d", service.Name, service.PingStatus))
			}
			fmt.Println(team.Name, servicesInfo)

		}
		fmt.Println()
	}
}
