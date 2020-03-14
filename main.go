package main

import (
	"github.com/gin-gonic/gin"
	"websocket/src/demo"
)



func main() {
	bindAddress := "localhost:2303"
	router := gin.Default()
	router.LoadHTMLGlob("src/templates/*")
	router.StaticFile("/favicon.ico", "src/assets/favicon.ico")
	router.GET("/ping/:code", demo.Ping)
	router.GET("/ping2/:code", demo.Ping2)
	router.GET("/echo", demo.Echo)
	router.GET("/write/:code/:title/:body", demo.Write)
	router.GET("/", demo.WebIndex)
	router.Run(bindAddress)
}
