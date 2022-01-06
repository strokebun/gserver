package iface

// @Description: 客户端请求封装，封装连接信息和请求数据
// @Author: StrokeBun
// @Date: 2022/1/6 19:16
type IRequest interface {
	// 获取连接
	GetConnection() IConnection
	// 获取数据
	GetData() []byte
}