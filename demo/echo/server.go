package main

import (
	"fmt"
	"github.com/strokebun/gserver/iface"
	"github.com/strokebun/gserver/server"
)

// @Description: 测试服务器
// @Author: StrokeBun
// @Date: 2022/1/6 16:17
type TestRouter struct {
	server.BaseRouter
}

func (r *TestRouter) Handle(request iface.IRequest)  {
	conn := request.GetConnection()
	fmt.Println("receive from client, msgId = ", request.GetMsgId(), ", data = ", string(request.GetData()))
	conn.SendMsg(1, request.GetData())
}

func main() {
	server := server.NewServer()
	router := &TestRouter{struct{}{}}
	server.AddRouter(1, router)
	server.Serve()
}