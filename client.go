package main

import (
	"fmt"
	"websocket/src/ws/client"
)

func main(){
	ws := client.Hup{}
	ws.Url = "ws://localhost:2303/ping/123"
	ws.Func = func(mt int,message string) {
		fmt.Println(message)

	}
	ws.WsClient()
}