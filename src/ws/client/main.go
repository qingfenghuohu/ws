package client

import (
	"github.com/gorilla/websocket"
	"log"
)

type Hup struct {
	Url  string
	Func func(mt int,message string)
	Ws   *websocket.Conn
}

func (c *Hup) WsClient() {
	ws, _, err := websocket.DefaultDialer.Dial(c.Url, nil)
	c.Ws = ws
	if err != nil {
		log.Fatal("WsClient:", err)
	}
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) != "" {
			c.Func(mt,string(message))
		}
	}
}
func (c *Hup) WriteMessage(Mt int,Mess string) {
	err := c.Ws.WriteMessage(Mt,[]byte(Mess))
	if err != nil{
		log.Fatal("WsClient WriteMessage:", err)
	}
}

