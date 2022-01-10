package core

import (
	"errors"
	"fmt"
	"github.com/strokebun/gserver/conf"
	"github.com/strokebun/gserver/iface"
	"io"
	"net"
)

// @Description: 连接模块
// @Author: StrokeBun
// @Date: 2022/1/6 18:30
type Connection struct {
	server iface.IServer
	// 当前连接的tcp socket
	conn *net.TCPConn
	// 连接id
	connId uint32
	// 连接是否关闭
	isClosed bool
	// 告知当前连接已经停止的channel
	exitChan chan bool
	// 同步读写操作
	msgChan chan []byte
	// 当前连接对应的路由模块
	msgHandler iface.IMessageHandler
}

// 创建连接
func NewConnection(server iface.IServer, conn *net.TCPConn, connID uint32, msgHandler iface.IMessageHandler) *Connection {
	c := &Connection{
		server:     server,
		conn:       conn,
		connId:     connID,
		msgHandler: msgHandler,
		isClosed:   false,
		exitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
	}
	server.GetConnManager().Add(c)
	return c
}

// 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println("connId =", c.connId, ",reader has exited")
	defer c.Stop()

	for {
		dataPack := NewDataPack()
		header := make([]byte, dataPack.GetHeaderLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), header); err != nil {
			fmt.Println("read msg header err", err)
			break
		}

		msg, err := dataPack.Unpack(header)
		if err != nil {
			fmt.Println("unpack err,", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read data err,", err)
				break
			}
		}
		msg.SetData(data)

		request := &Request{
			conn: c,
			msg: msg,
		}

		if conf.GlobalObject.WorkPoolSize > 0 {
			c.msgHandler.SendMsgToTaskQueue(request)
		} else {
			go c.msgHandler.DoMessageHandler(request)
		}

	}
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.conn.Write(data); err != nil {
				fmt.Println("Send Data error,", err, ", Conn Writer exit")
				return
			}
		case <-c.exitChan:
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("connection start.. ConnID =", c.connId)
	go c.StartReader()
	go c.StartWriter()
	// 调用连接创建的hook
	c.server.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	fmt.Println("connection stop.. ConnID =", c.connId)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.conn.Close()
	// 从连接管理器删除该连接
	c.server.GetConnManager().Remove(c)
	// 调用连接停止的hook
	c.server.CallOnConnStop(c)
	close(c.exitChan)
	close(c.msgChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetConnectionId() uint32 {
	return c.connId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
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

	c.msgChan <- binaryMsg
	return nil
 }