package server

import (
	"fmt"
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
}

func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      6023,
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

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept TCP err", err)
				continue
			}

			go func() {
				for {
					data := make([]byte, 512)
					count, err := conn.Read(data)
					if err != nil {
						fmt.Println("receive buf err ", err)
						continue
					}
					if _, err := conn.Write(data[:count]); err != nil {
						fmt.Println("write back buffer err ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) Stop() {

}