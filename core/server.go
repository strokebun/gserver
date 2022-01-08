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
	IPVersion string
	// 服务绑定的IP地址
	IP string
	// 服务绑定的端口
	Port int
	// 路由模块
	MsgHandler iface.IMessageHandler
	// 连接管理模块
	ConnManager iface.IConnectionManager

	// 连接创建时的hook函数
	OnConnStart func(conn iface.IConnection)
	// 连接停止时的hook函数
	OnConnStop func(conn iface.IConnection)
}

func NewServer() *Server {
	return &Server{
		Name:      conf.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        conf.GlobalObject.Host,
		Port:      conf.GlobalObject.Port,
		MsgHandler: NewMessageHandler(),
		ConnManager: NewConnectionManager(),
	}
}

func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s, listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)
	go func() {
		s.MsgHandler.StartWorkerPool()
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp address error: ", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err: ", err)
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
			if s.ConnManager.ConnNum() >= conf.GlobalObject.MaxConn {
				conn.Close()
				continue
			}

			connId++
			dealConn := NewConnection(s, conn, connId, s.MsgHandler)
			go dealConn.Start()
		}
	}()
}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) Stop() {
	fmt.Println("[STOP] gserver stop, name:", s.Name)
	s.ConnManager.Clear()
}

func (s *Server) AddRouter(msgId uint32, router iface.IRouter)  {
	s.MsgHandler.AddRouter(msgId, router)
}

func (s *Server) GetConnManager() iface.IConnectionManager {
	return s.ConnManager
}

func (s *Server) SetOnConnStart(hook func(connection iface.IConnection)) {
	s.OnConnStart = hook
}

func (s *Server) SetOnConnStop(hook func(connection iface.IConnection)) {
	s.OnConnStop = hook
}

// 调用连接创建时的hook
func (s *Server) CallOnConnStart(connection iface.IConnection) {
	if s.OnConnStart != nil {
		s.OnConnStart(connection)
	}
}

// 调用连接停止时的hook
func (s *Server) CallOnConnStop(connection iface.IConnection) {
	if s.OnConnStart != nil {
		s.OnConnStop(connection)
	}
}