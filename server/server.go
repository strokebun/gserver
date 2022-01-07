package server

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
}

func NewServer() *Server {
	return &Server{
		Name:      conf.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        conf.GlobalObject.Host,
		Port:      conf.GlobalObject.Port,
		MsgHandler: NewMessageHandler(),
	}
}

func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s, listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)
	go func() {
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
		fmt.Println("[SUCCESS] start server ", s.Name, " successfully. Listening...")

		var connId uint32 = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept TCP err", err)
				continue
			}

			connId++
			dealConn := NewConnection(conn, connId, s.MsgHandler)
			go dealConn.Start()
		}
	}()
}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) Stop() {

}

func (s *Server) AddRouter(msgId uint32, router iface.IRouter)  {
	s.MsgHandler.AddRouter(msgId, router)
}
