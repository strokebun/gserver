package iface

import "net"

// @Description: 连接模块抽象接口
// @Author: StrokeBun
// @Date: 2022/1/6 18:30
type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取当前连接绑定的socket connection
	GetTCPConnection() *net.TCPConn
	// 获取连接id
	GetConnectionId() uint32
	// 获取客户端状态
	RemoteAddr() net.Addr
	// 发送数据
	Send([]byte)
}

// 当前连接处理函数
type Handler func(*net.TCPConn, []byte, int) error
