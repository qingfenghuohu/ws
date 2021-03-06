package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/gorilla/websocket"
	"net/http"
)

//var upGrader = Upgrader{
//	//CheckOrigin: func (r *http.Request) bool {
//	//	return true
//	//},
//}

type Message struct {
	Id   string
	Type int
	Body string
}
type Hup struct {
	WsList     map[string]map[int]*Conn
	wsListMark map[string]int
	Read       map[string]chan Message
	UpGrader   Upgrader
}

func NewWs() Hup {
	WsHup := Hup{}
	WsHup.WsList = map[string]map[int]*Conn{}
	WsHup.wsListMark = map[string]int{}
	WsHup.UpGrader = Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	WsHup.Read = map[string]chan Message{}
	return WsHup
}

func (h *Hup) SetUpGrader(Ug Upgrader) {
	h.UpGrader = Ug
}

func (h *Hup) WsHandler(Id string, c *gin.Context, callback func(conn *Conn, Id string, num int), ReadMessage func(Id string, c *Conn, h *Hup)) {
	ws, err := h.UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	markNum := h.wsListMark[Id] + 1
	h.wsListMark[Id]++
	if len(h.WsList[Id]) == 0 {
		h.WsList[Id] = map[int]*Conn{}
	}
	h.WsList[Id][markNum] = ws
	go func() {
		callback(ws, Id, markNum)
	}()
	defer func() {
		h.Close(ws, Id, markNum)
	}()
	h.ReadMessage(Id, ws, ReadMessage)
}
func (h *Hup) Close(conn *Conn, Id string, MarkNum int) {
	fmt.Println("断开连接")
	delete(h.WsList[Id], MarkNum)
	conn.Close()
	if len(h.WsList[Id]) == 0 {
		h.wsListMark[Id] = 0
		delete(h.WsList, Id)
	}
}

func (h *Hup) WriteMessage(Mess Message) {
	if ok := h.WsList[Mess.Id]; len(ok) > 0 {
		for _, v := range h.WsList[Mess.Id] {
			err := v.WriteMessage(Mess.Type, []byte(Mess.Body))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (h *Hup) AllWriteMessage(t int, content []byte) {
	for _, vv := range h.WsList {
		for _, v := range vv {
			err := v.WriteMessage(t, content)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (h *Hup) ReadMessage(Id string, c *Conn, callback func(Id string, c *Conn, h *Hup)) {
	for {
		callback(Id, c, h)
	}
}
