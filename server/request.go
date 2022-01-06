package server

import "github.com/strokebun/gserver/iface"

// @Description: 客户端请求
// @Author: StrokeBun
// @Date: 2022/1/6 19:17
type Request struct {
	// 和客户端建立的连接
	conn iface.IConnection
	// 客户端请求数据
	data []byte
}

func (r *Request) GetConnection() iface.IConnection  {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}