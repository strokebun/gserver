package core

import (
	"fmt"
	"github.com/strokebun/gserver/conf"
	"github.com/strokebun/gserver/iface"
	"net"
)

// @Description: Server服务类，实现IServer接口
// @Author: StrokeBun
// @Date: 2022/1/6 16:11
type Server struct {
	// 服务器的名称
	Name string
	// tcp4 or other
	ipVersion string
	// 服务绑定的IP地址
	ip string
	// 服务绑定的端口
	port int
	// 路由模块
	msgHandler iface.IMessageHandler
	// 连接管理模块
	connManager iface.IConnectionManager

	// 连接创建时的hook函数
	onConnStart func(conn iface.IConnection)
	// 连接停止时的hook函数
	onConnStop func(conn iface.IConnection)
}

func NewServer() *Server {
	return &Server{
		Name:      conf.GlobalObject.Name,
		ipVersion: "tcp4",
		ip:        conf.GlobalObject.Host,
		port:      conf.GlobalObject.Port,
		msgHandler: NewMessageHandler(),
		connManager: NewConnectionManager(),
	}
}

func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s, listenner at IP: %s, Port %d is starting\n", s.Name, s.ip, s.port)
	go func() {
		s.msgHandler.StartWorkerPool()
		addr, err := net.ResolveTCPAddr(s.ipVersion, fmt.Sprintf("%s:%d", s.ip, s.port))
		if err != nil {
			fmt.Println("resolve tcp address error:", err)
			return
		}
		listener, err := net.ListenTCP(s.ipVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.ipVersion, " err: ", err)
			return
		}
		fmt.Println("[SUCCESS] start server", s.Name, "successfully. Listening...")

		var connId uint32 = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept TCP err", err)
				continue
			}
			// 超过最大连接数，关闭该连接
			if s.connManager.ConnNum() >= conf.GlobalObject.MaxConn {
				conn.Close()
				continue
			}

			connId++
			dealConn := NewConnection(s, conn, connId, s.msgHandler)
			go dealConn.Start()
		}
	}()
}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) Stop() {
	fmt.Println("[STOP] server stop, name:", s.Name)
	s.connManager.Clear()
}

func (s *Server) AddRouter(msgId uint32, router iface.IRouter)  {
	s.msgHandler.AddRouter(msgId, router)
}

func (s *Server) GetConnManager() iface.IConnectionManager {
	return s.connManager
}

func (s *Server) SetOnConnStart(hook func(connection iface.IConnection)) {
	s.onConnStart = hook
}

func (s *Server) SetOnConnStop(hook func(connection iface.IConnection)) {
	s.onConnStop = hook
}

// 调用连接创建时的hook
func (s *Server) CallOnConnStart(connection iface.IConnection) {
	if s.onConnStart != nil {
		s.onConnStart(connection)
	}
}

// 调用连接停止时的hook
func (s *Server) CallOnConnStop(connection iface.IConnection) {
	if s.onConnStart != nil {
		s.onConnStop(connection)
	}
}