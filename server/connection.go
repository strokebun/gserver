package server

import (
	"fmt"
	"github.com/strokebun/gserver/iface"
	"net"
)

// @Description: 连接模块
// @Author: StrokeBun
// @Date: 2022/1/6 18:30
type Connection struct {
	// 当前连接的tcp socket
	Conn *net.TCPConn
	// 连接id
	ConnID uint32
	// 连接是否关闭
	isClosed bool
	// 当前连接的处理函数
	handleAPI iface.Handler
	// 告知当前连接已经停止的channel
	ExitChan chan bool
}

// 创建连接
func NewConnection(conn *net.TCPConn, connID uint32, handlerAPI iface.Handler) *Connection {
	return &Connection{
		Conn:        conn,
		ConnID:      connID,
		isClosed:    false,
		handleAPI: handlerAPI,
		ExitChan: make(chan bool, 1),
	}
}

// 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("reader is running..")
	defer fmt.Println("connId = ", c.ConnID, " reader is exit")
	defer c.Stop()

	for {
		data := make([]byte, 512)
		count, err := c.Conn.Read(data)
		if err != nil {
			fmt.Println("receive buf err ", err)
			continue
		}
		if err := c.handleAPI(c.Conn, data, count); err != nil {
			fmt.Println("connId ", c.ConnID, " err: ", err)
			break
		}

	}
}

func (c *Connection) Start() {
	fmt.Println("connection start.. ConnID= ", c.ConnID)
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("connection stop.. ConnID = ", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnectionID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c* Connection) Send(data []byte) {

 }