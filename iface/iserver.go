package iface

// @Description: Server接口，定义Server的服务
// @Author: StrokeBun
// @Date: 2022/1/6 16:08
type IServer interface {
	// 启动服务器
	Start()
	// 开启业务服务
	Serve()
	// 关闭服务器
	Stop()
	// 添加路由, 供客户端的连接调用
	AddRouter(msgId uint32, router IRouter)
	// 获得连接管理器
	GetConnManager() IConnectionManager

	// 设置连接创建时的hook
	SetOnConnStart(func(IConnection))
	// 设置连接断开时的hook
	SetOnConnStop(func(IConnection))
	// 调用连接创建时的hook
	CallOnConnStart(connection IConnection)
	// 调用连接停止时的hook
	CallOnConnStop(connection IConnection)
}
