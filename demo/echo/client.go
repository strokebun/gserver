package main

import (
	"fmt"
	"github.com/strokebun/gserver/core"
	"io"
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
		clientMsg := core.NewMessage(1, []byte("echo test..."))
		dataPack := core.NewDataPack()
		binaryMsg, err := dataPack.Pack(clientMsg)
		if err != nil {
			fmt.Println("pack error ", err)
			return
		}

		_, err = conn.Write(binaryMsg)
		if err != nil {
			fmt.Println("conn write err", err)
			return
		}

		// 读取头部
		header := make([]byte, dataPack.GetHeaderLen())
		if _, err := io.ReadFull(conn, header); err != nil {
			fmt.Println("read header err", err)
			break
		}
		headMsg, err := dataPack.Unpack(header)
		if err != nil {
			fmt.Println("client unpack err", err)
			break
		}
		if headMsg.GetMsgLen() > 0 {
			// 读入数据
			serverMsg := headMsg.(*core.Message)
			serverMsg.Data = make([]byte, headMsg.GetMsgLen())
			if _, err := io.ReadFull(conn, serverMsg.Data); err != nil {
				fmt.Println("client read msg err", err)
				return
			}
			fmt.Printf("receive core call back: msg id = %d, data = %s\n", serverMsg.Id, serverMsg.Data)
		}

		time.Sleep(time.Second)
	}
}