package core

import (
	"fmt"
	"github.com/strokebun/gserver/conf"
	"github.com/strokebun/gserver/iface"
)

// @Description: 消息处理模块实现
// @Author: StrokeBun
// @Date: 2022/1/7 15:27
type MessageHandler struct {
	// msgId以及对应的处理器
	apis map[uint32]iface.IRouter
	// 工作池大小
	workPoolSize uint32
	// 任务队列
	taskQueue []chan iface.IRequest
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		apis: make(map[uint32]iface.IRouter),
		workPoolSize: conf.GlobalObject.WorkPoolSize,
		taskQueue: make([]chan iface.IRequest, conf.GlobalObject.WorkPoolSize),
	}
}

func (mh *MessageHandler) DoMessageHandler(request iface.IRequest) {
	router, ok := mh.apis[request.GetMsgId()]
	if !ok {
		fmt.Println("[WARNING] api msgId =", request.GetMsgId(), "miss")
		return
	}
	router.PreHandle(request)
	router.Handle(request)
	router.AfterHandle(request)
}

func (mh *MessageHandler) AddRouter(msgId uint32, router iface.IRouter) {
	if _, ok := mh.apis[msgId]; ok {
		fmt.Println("[WARNING] msgId has existed...")
	}
	mh.apis[msgId] = router
	fmt.Println("add api msgId =", msgId, "success")
}

// 启动一个Worker工作流程
func (mh *MessageHandler) StartOneWorker(workerID int, taskQueue chan iface.IRequest) {
	fmt.Println("Worker ID =", workerID, "started.")
	for {
		select {
		case request := <-taskQueue:
			mh.DoMessageHandler(request)
		}
	}
}

// 启动worker工作池
func (mh *MessageHandler) StartWorkerPool() {
	for i := 0; i < int(mh.workPoolSize); i++ {
		mh.taskQueue[i] = make(chan iface.IRequest, conf.GlobalObject.MaxWorkTaskLen)
		go mh.StartOneWorker(i, mh.taskQueue[i])
	}
}

// 将消息交给任务队列,由worker进行处理
func (mh *MessageHandler) SendMsgToTaskQueue(request iface.IRequest) {
	// 根据connID轮询分配任务
	workerID := request.GetConnection().GetConnectionId() % mh.workPoolSize
	mh.taskQueue[workerID] <- request
}

