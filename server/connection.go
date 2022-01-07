package server

import (
	"errors"
	"fmt"
	"github.com/strokebun/gserver/iface"
	"io"
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
	// 告知当前连接已经停止的channel
	ExitChan chan bool
	// 当前连接对应的路由模块
	Router iface.IRouter
}

// 创建连接
func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	return &Connection{
		Conn:        conn,
		ConnID:      connID,
		Router: router,
		isClosed:    false,
		ExitChan: make(chan bool, 1),
	}
}

// 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("reader is running..")
	defer fmt.Println("connId = ", c.ConnID, " reader is exit")
	defer c.Stop()

	for {
		dataPack := NewDataPack()
		header := make([]byte, dataPack.GetHeaderLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), header); err != nil {
			fmt.Println("read msg header err ", err)
			break
		}

		msg, err := dataPack.Unpack(header)
		if err != nil {
			fmt.Println("unpack err ", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read data ", err)
				break
			}
		}
		msg.SetData(data)

		request := Request{
			conn: c,
			msg: msg,
		}
		go func(req iface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.AfterHandle(req)
		}(&request)
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

func (c *Connection) GetConnectionId() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c* Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection has closed")
	}

	msg := NewMessage(msgId, data)
	dataPack := NewDataPack()
	binaryMsg, err := dataPack.Pack(msg)
	if err != nil {
		fmt.Println("pack error, msg id =", msgId)
		return errors.New("pack message error")
	}

	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("write msg error, msg id =", msgId)
		return errors.New("conn write error")
	}
	return nil
 }