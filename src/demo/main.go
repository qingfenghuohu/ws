package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"websocket/src/ws/server"
	. "github.com/gorilla/websocket"
	"github.com/gorilla/websocket"


)
var Ws server.Hup

func init() {
	Ws = server.NewWs()
}
func Write(c *gin.Context) {
	code := c.Param("code")
	title := c.Param("title")
	body := c.Param("body")

	Ws.WriteMessage(server.Message{code,1,title+body})
}

//webSocket请求ping 返回pong
//noinspection ALL
func Ping(c *gin.Context) {
	code := c.Param("code")
	Ws.WsHandler(code,c)

}
func WebIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
var upGrader = Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}
func Ping2(c *gin.Context){
	//code := c.Param("code")
	//升级get请求为webSocket协议
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func() {
		fmt.Println("执行关闭")
		ws.Close()
	}()

	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		//写入ws数据
		err = ws.WriteMessage(mt,message)
		if err != nil {
			break
		}
	}
}
var upgrader = websocket.Upgrader{} // use default options

func Echo(c *gin.Context){
	var ws, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}