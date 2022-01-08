package iface

// @Description: 连接管理模块
// @Author: StrokeBun
// @Date: 2022/1/8 17:08
type IConnectionManager interface {
	// 添加连接
	Add(conn IConnection)
	// 删除连接
	Remove(conn IConnection)
	// 根据ConnID获取连接
	Get(connID uint32) (IConnection, error)
	// 获取当前连接数量
	ConnNum() int
	// 删除并停止所有连接
	Clear()
}
