package main

import (
	"fmt"
	"github.com/strokebun/gserver/core"
	"github.com/strokebun/gserver/iface"
)

// @Description: 测试服务器
// @Author: StrokeBun
// @Date: 2022/1/6 16:17
type TestRouter struct {
	core.BaseRouter
}

func (r *TestRouter) Handle(request iface.IRequest)  {
	conn := request.GetConnection()
	fmt.Println("receive from client, msgId =", request.GetMsgId(), ", data =", string(request.GetData()))
	conn.SendMsg(1, request.GetData())
}

func connStartHook(connection iface.IConnection) {
	fmt.Println("conn start hook call, connId =", connection.GetConnectionId())
}

func connStopHook(connection iface.IConnection) {
	fmt.Println("conn stop hook call, connId =", connection.GetConnectionId())
}

func main() {
	server := core.NewServer()
	router := &TestRouter{struct{}{}}
	server.AddRouter(1, router)
	server.SetOnConnStart(connStartHook)
	server.SetOnConnStop(connStopHook)
	server.Serve()
}