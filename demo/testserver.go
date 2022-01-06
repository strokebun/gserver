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
	data := request.GetData()
	fmt.Println("call back to client")
	conn.GetTCPConnection().Write(data)
}

func main() {
	server := server.NewServer("test server")
	router := &TestRouter{struct{}{}}
	server.AddRouter(router)
	server.Serve()
}