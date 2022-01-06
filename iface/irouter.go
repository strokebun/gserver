package iface

// @Description: 路由模块接口
// @Author: StrokeBun
// @Date: 2022/1/6 19:23
type IRouter interface {
	// 处理业务之前的hook
	PreHandle(request IRequest)
	// 处理业务
	Handle(request IRequest)
	// 处理业务之后的hook
	AfterHandle(request IRequest)
}