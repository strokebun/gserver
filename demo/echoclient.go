package main

import (
	"fmt"
	"net"
	"time"
)

// @Description: echo客户端，用于测试
// @Author: StrokeBun
// @Date: 2022/1/6 16:17
func main() {
	fmt.Println("[START] echo client start")
	conn, err := net.Dial("tcp", "127.0.0.1:6023")
	if err != nil {
		fmt.Println("client start err ", err)
		return
	}
	for {
		_, err = conn.Write([]byte("echo test..."))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err ", err)
			return
		}
		fmt.Printf("receive server call back: %s\n", buf)
		time.Sleep(time.Second)
	}
}