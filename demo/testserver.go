package main

import "github.com/StrokeBun/gserver/server"

// @Description: 测试服务器
// @Author: StrokeBun
// @Date: 2022/1/6 16:17
func main() {
	server := server.NewServer("test server")
	server.Serve()
}