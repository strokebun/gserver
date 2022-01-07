package iface

// @Description: 消息处理模块，管理不同消息
// @Author: StrokeBun
// @Date: 2022/1/7 15:26
type IMessageHandler interface {
	// 处理消息
	DoMessageHandler(request IRequest)
	// 添加消息路由
	AddRouter(msgId uint32, router IRouter)
	// 启动worker工作池
	StartWorkerPool()
	// 将消息交给任务队列
	SendMsgToTaskQueue(request IRequest)
}
