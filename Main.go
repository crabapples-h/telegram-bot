package main

import (
	"github.com/gin-gonic/gin"
	"log"
	myApi "telegram-bot/api"
	_ "telegram-bot/dao"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
}

func getProxy(context *gin.Context) {
	name := context.Param("name")
	log.Printf("收到请求->获取代理IP,服务->[%s]\n", name)
	log.Printf("返回结果->获取代理IP,服务->[%s]\n", name)
}

func InitRouter() *gin.Engine {
	server := gin.Default()
	server.POST("/api/callback/update", myApi.CallBack)
	return server
}
func main() {
	server := InitRouter()
	err := server.Run(":9093")
	log.Println("服务启动完成:]")
	if err != nil {
		log.Printf("服务启动异常:[%s]\n", err.Error())
	}
}
