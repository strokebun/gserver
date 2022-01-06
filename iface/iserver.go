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
}
