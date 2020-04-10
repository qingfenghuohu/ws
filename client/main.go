package client

import (
	"github.com/gorilla/websocket"
	"log"
)

type Hup struct {
	Url string
	Ws  *websocket.Conn
}

func (c *Hup) WsClient(callback func(ws *websocket.Conn)) {
	ws, _, err := websocket.DefaultDialer.Dial(c.Url, nil)
	c.Ws = ws
	if err != nil {
		log.Fatal("WsClient:", err)
	}
	for {
		callback(ws)
	}
}
func (c *Hup) WriteMessage(Mt int, Mess string) {
	err := c.Ws.WriteMessage(Mt, []byte(Mess))
	if err != nil {
		log.Fatal("WsClient WriteMessage:", err)
	}
}
